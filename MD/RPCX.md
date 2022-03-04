# 注册中心

 用来实现服务发现和服务的元数据存储

# 服务注册

服务端提供者将服务的元数据信息注册到注册中心的过程。

元数据包括：

- 服务名
- 监听地址
- 监听协议
- 权重 吞吐率等

```go
func (s *Server) RegisterName(name string, rcvr interface{}, metadata string) error {
    //在注册中心做服务注册
    s.Plugins.DoRegister(name, rcvr, metadata)
    //通过反射，获取rcvr的Type，Value，Method等信息
    _, err := s.register(rcvr, name, true)
    return err
}
```



# 服务发现

客户端获取服务元数据的过程。

两种实现方式

- pull 
- push

![](https://pic2.zhimg.com/80/v2-a3b15c205dd3984b253b507614cc4461_1440w.jpg)

```go
type xClient struct {
    failMode     FailMode //失败策略：Failover，Failfast，Failtry，Failbackup
    selectMode   SelectMode //调用策略：RandomSelect，RoundRobin，WeightedRoundRobin，WeightedICMP等
    ……
    servers   map[string]string  //实时的注册服务信息
    discovery ServiceDiscovery  //ZookeeperDiscovery
    ……
}
```



# 远程调用

这个过程中需要涉及到服务治理，如果序列化和反序列化。

## 服务治理

### 远程调用失败策略

- Failover
- Failfast
- Failtry
- Backup

### 路由算法

- 随机
- 轮询
- 权重
- 网络质量
- 一致性哈希
- 地理位置geo

Call方法做了什么事情？

```go

// Call invokes the named function, waits for it to complete, and returns its error status.
// It handles errors base on FailMode.
func (c *xClient) Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error {
    //如果连接已经中断
    if c.isShutdown {
        return ErrXClientShutdown
    }
​
    if c.auth != "" {
        metadata := ctx.Value(share.ReqMetaDataKey)
        if metadata == nil {
            metadata = map[string]string{}
            ctx = context.WithValue(ctx, share.ReqMetaDataKey, metadata)
        }
        m := metadata.(map[string]string)
        m[share.AuthKey] = c.auth
    }
​
    var err error
    //1、路由算法
    k, client, err := c.selectClient(ctx, c.servicePath, serviceMethod, args)
    if err != nil {
        //“在路由发生错误的时候”，如果是最快失败策略，则直接返回，不做任何重试
        if c.failMode == Failfast {
            return err
        }
    }
​
    var e error
    //2、失败策略
    switch c.failMode {
    case Failtry:
        ……
        //3、远程调用
        err = c.wrapCall(ctx, client, serviceMethod, args, reply)
        ……
    case Failover:
        ……
        //3、远程调用
        err = c.wrapCall(ctx, client, serviceMethod, args, reply)
        ……
    case Failbackup:
        ……
    default: //Failfast
        ……
        //3、远程调用
        err = c.wrapCall(ctx, client, serviceMethod, args, reply)
        ……
    }
}
```

进入路由算法

```go
// selects a client from candidates base on c.selectMode
func (c *xClient) selectClient(ctx context.Context, servicePath, serviceMethod string, args interface{}) (string, RPCClient, error) {
    //注意，这里是有锁的，路由是需要同步的
    c.mu.Lock()
    //获取路由算法
    var fn = c.selector.Select
    //用到装饰器模式，用Plugins来装饰路由器
    if c.Plugins != nil {
        fn = c.Plugins.DoWrapSelect(fn)
    }
    //选择服务节点
    k := fn(ctx, servicePath, serviceMethod, args)
    c.mu.Unlock()
    if k == "" {
        return "", nil, ErrXClientNoServer
    }
    //获取一个连接
    client, err := c.getCachedClient(k)
    return k, client, err
}
```

```go
func newSelector(selectMode SelectMode, servers map[string]string) Selector {
	switch selectMode {
    // 随机
	case RandomSelect:
		return newRandomSelector(servers)
    // 轮询
	case RoundRobin:
		return newRoundRobinSelector(servers)
    // 权重
	case WeightedRoundRobin:
		return newWeightedRoundRobinSelector(servers)
    // ping 
	case WeightedICMP:
		return newWeightedICMPSelector(servers)
    // 一致性hash
	case ConsistentHash:
		return newConsistentHashSelector(servers)
    // 用户选择
	case SelectByUser:
		return nil
	default:
		return newRandomSelector(servers)
	}
}
```





# 序列化

rpcx支持的序列化方式

```go
作者：修华师
链接：https://zhuanlan.zhihu.com/p/144374225
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

var (
    // Codecs are codecs supported by rpcx. You can add customized codecs in Codecs.
    Codecs = map[protocol.SerializeType]codec.Codec{
        //用户自己实现与[]byte直接的转换
        protocol.SerializeNone: &codec.ByteCodec{},
        //json
        protocol.JSON:          &codec.JSONCodec{},
        //protobuf:github.com/gogo/protobuf 使用的并非官方库
        protocol.ProtoBuffer:   &codec.PBCodec{},
        //msgpack
        protocol.MsgPack:       &codec.MsgpackCodec{},
        //thrift
        protocol.Thrift:        &codec.ThriftCodec{},
    }
)
```

默认使用MsgPack

```go
var DefaultOption = Option{
    Retries:        3,
    RPCPath:        share.DefaultRPCPath,
    ConnectTimeout: 10 * time.Second,
    Breaker:        CircuitBreaker,
    SerializeType:  protocol.MsgPack,//默认MsgPack
    CompressType:   protocol.None,
}

```





kit：一个微服务的基础库（框架）
service：业务代码 + kit 依赖 + 第三方依赖组成的业务微服务
RPC + message queue：轻量级通讯



移动端 -> API Gateway -> BFF -> Microservices，在 FE Web业务中，BFF 可以是 Node.js 来做服务端渲染（SSR，Server-Side Rendering），注意这里忽略了上游的 CDN、4/7层负载均衡（ELB）。