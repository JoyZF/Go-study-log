Redis排查文档

###### 延迟基线测量

如何判断Redis真的变慢了呢？

当你发现Redis运行时的延迟时基线性能的2倍以上时就可以判断Redis性能变慢。

redis-cli提供了-intrinsic-latency选项，用来监测和统计测试期间内的最大延迟，这个延迟可以作为redis的基线性能。

```shell
redis-cli --latency -h `host` -p `port`
```

注意需要用服务端运行，避免网络对基线性能的影响。

### 慢指令监控

如何判断是否是慢指令呢？

操作复杂度为o（N）的都认为是慢指令。

比如HGETALL、SMEMBERS以及SORT、LREM、SUNION

如何监控呢？

- 使用Redis慢日志功能查询
- latency-monitor 延迟监控工具

### 慢日志功能

redis中的slowlog 命令可以让我们快速定位到那些超出指定执行时间的慢命令，默认情况下命令是记录10ms。

可以在redis-cli中输入命令配置：

```shell
redis-cli CONFIG SET slowlog-log-slower-than 6000
```

也可以在Redis.Config配置文件中配置。

查看所有慢命令可以使用slowlog get命令查看，返回结果的第三个字段以微妙单位显示命令的执行时间。

如果要查询最后2个慢命令，可以输入slowlog get 2即可。

### Latency Monitoring

Latency Monitoring 用于以秒为粒度监控各种事件发生的频率。

```shell
CONFIG SET latency-monitoring 9 
```



### 如何解决Redis变慢的问题？

#### 网络通信导致的延迟。

#### 慢指令导致的延迟

- 在集群中，将聚合运算等o（n）的操作放在slave上。
- 使用高效的命令代替

##### FORK生成RDB导致延迟

生成RDB快照，Redis必须fork后台进程，fork操作会导致主进程阻塞。可以使用basave

#### 内存大页

常规的内存页是按照 4 KB 来分配，Linux 内核从 2.6.38 开始支持内存大页机制，该机制支持 2MB 大小的内存页分配。

Redis 使用了 fork 生成 [RDB 做持久化提供了数据可靠性保证](https://mp.weixin.qq.com/s?__biz=MzkzMDI1NjcyOQ==&mid=2247487758&idx=1&sn=beb5918bb61948b2920907f54510311f&scene=21#wechat_redirect)。

当生成 RDB 快照的过程中，Redis 采用**[写时复制](https://mp.weixin.qq.com/s?__biz=MzkzMDI1NjcyOQ==&mid=2247487758&idx=1&sn=beb5918bb61948b2920907f54510311f&scene=21#wechat_redirect)**技术使得主线程依然可以接收客户端的写请求。

也就是当数据被修改的时候，Redis 会复制一份这个数据，再进行修改。

采用了内存大页，生成 RDB 期间，即使客户端修改的数据只有 50B 的数据，Redis 需要复制 2MB 的大页。当写的指令比较多的时候就会导致大量的拷贝，导致性能变慢。

可以使用指令禁用Linux 内存大页

```shell
echo never > /sys/kernel/mm/transparent_hugepage/enabled
```

#### swap 操作系统分页

当物理内存（内存条）不够用的时候，将部分内存上的数据交换到 swap 空间上，以便让系统不会因内存不够用而导致 oom 或者更致命的情况出现。

当某进程向 OS 请求内存发现不足时，OS 会把内存中暂时不用的数据交换出去，放在 SWAP 分区中，这个过程称为 SWAP OUT。

当某进程又需要这些数据且 OS 发现还有空闲物理内存时，又会把 SWAP 分区中的数据交换回物理内存中，这个过程称为 SWAP IN。

**内存 swap 是操作系统里将内存数据在内存和磁盘间来回换入和换出的机制，涉及到磁盘的读写。**



#### expires 淘汰过期数据

Redis 有两种方式淘汰过期数据

- 惰性删除：当接受请求的时候发现key已经过期，才执行删除
- 定时删除：每100ms删除一些过期key

解决办法：

如果一批 key 的确是同时过期，可以在 `EXPIREAT` 和 `EXPIRE` 的过期时间参数上，**加上一个一定大小范围内的随机数**，这样，既保证了 key 在一个邻近时间范围内被删除，又避免了同时过期造成的压力。

#### bigkey

通常我们会将含有较大数据或含有大量成员、列表数的 Key 称之为大 Key，下面我们将用几个实际的例子对大 Key 的特征进行描述：

- 一个 STRING 类型的 Key，它的值为 5MB（数据过大）
- 一个 LIST 类型的 Key，它的列表数量为 10000 个（列表数量过多）
- 一个 ZSET 类型的 Key，它的成员数量为 10000 个（成员数量过多）
- 一个 HASH 格式的 Key，它的成员数量虽然只有 1000 个但这些成员的 value 总大小为 10MB（成员体积过大）

bigkey 带来问题如下：

1. Redis 内存不断变大引发 OOM，或者达到 maxmemory 设 置值引发写阻塞或重要 Key 被逐出；
2. Redis Cluster 中的某个 node 内存远超其余 node，但因 Redis Cluster 的数据迁移最小粒度为 Key 而无法将 node 上的内存均衡化；
3. bigkey 的读请求占用过大带宽，自身变慢的同时影响到该服务器上的其它服务；
4. 删除一个 bigkey 造成主库较长时间的阻塞并引发同步中断或主从切换；

### 总结

如下检查清单，帮助你在遇到 Redis 性能变慢的时候能高效解决问题。

1. 获取当前 Redis 的基线性能；
2. 开启慢指令监控，定位慢指令导致的问题；
3. 找到慢指令，使用 scan 的方式；
4. 将实例的数据大小控制在 2-4GB，避免主从复制加载过大 RDB 文件而阻塞；
5. 禁用内存大页，采用了内存大页，生成 RDB 期间，即使客户端修改的数据只有 50B 的数据，Redis 需要复制 2MB 的大页。当写的指令比较多的时候就会导致大量的拷贝，导致性能变慢。
6. Redis 使用的内存是否过大导致 swap；
7. AOF 配置是否合理，可以将配置项 no-appendfsync-on-rewrite 设置为 yes，避免 AOF 重写和 fsync 竞争磁盘 IO 资源，导致 Redis 延迟增加。
8. bigkey 会带来一系列问题，我们需要进行拆分防止出现 bigkey，并通过 UNLINK 异步删除。

