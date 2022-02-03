# 项目的主干

```
/cmd
```

每个项目的应用程序目录应该和可执行文件的名称相匹配。比如/cmd/mkt.go，/cmd/mkt

不要在cmd中放太多代码，如果想倒入并在其他项目中使用，那么应该用/pkg，如果不想被其他人应用，那么应该在/internal中

```
/internal
```

私有应用程序和库代码，不能被其他项目引用。

```
/pkg
```

外部应用程序可以使用的库代码

# Kit Project Layout

每个公司应当为不同的微服务简历一个统一的kit工具包项目和app项目

kit项目必须具备的特点

- 统一
- 标准库方式布局
- 高度抽离
- 支持插件

# Service Application Project Layout

```
/api
```

API协议定义目录，用于放protobuf 文以及生成的go文件。

```
/configs
```

配置文件模版或默认配置

```
/test
```

额外的外部测试应用程序和测试数据
*不应该包含/src 因为这会和Go用于其工作空间的src目录相混搅*

## 微服务中app服务类型分为4类

- interface 对外的BFF服务，接收来自用户的请求
- service 对内的微服务，仅接受来自内部其他服务或者网关的请求。
- admin 区别于servive，更多的是面向运营侧的服务
- job 流式任务处理的服务，上游一般依赖于message broker
- task 定时任务



# Wire

google wire用于生成静态代码，时间依赖注入的思路DI

# gRPC

- 多语言
- 轻量级、高性能
- 可插拔
- IDL 基于文件定义服务，通过proto3 工具生成制定语言的数据接口、服务端接口以及客户端Stub
- 移动端：基于标准的HTTP2设计，支持双向流、消息头压缩，单TCP的多路复用、服务端推送等特性。

*不要过早关注性能问题，先标准化*

为了统一检索和规范API，可以建立一个统一的api仓库，整合所有对内对外的API.

# Test

## Unittest

## API



# VO、DTO、DO、PO的概念、区别和用处

- VO view object，视图对象，用于展示层
- DTO data transfer object ，数据传输对象，用于展示层和服务层之间的数据传输对象
- DO domain object 领域对象，从现实世界中抽象出来的有形或无形的业务实体
- PO Persisten object 持久化对象，跟持久层的数据结构形成一一对应的映射关系

- 用户发出请求，表单的数据在展示层被匹配为VO
- - 展示层把VO转化为服务层对应方法所要求的DTO,传送给服务层
  - - 服务层搜想根据DTO的数据构造一个DO ，调用DO的业务方法完成具体的业务
    - - 服务层把DO转化为持久层对应的PO，调用持久层的持久方法

![image-20211027223021502](/Users/joy/Library/Application Support/typora-user-images/image-20211027223021502.png)

参考文章：[浅析VO、DTO、DO、PO的概念、区别和用处](https://www.cnblogs.com/zxf330301/p/6534643.html)](https://www.cnblogs.com/zxf330301/p/6534643.html)

# DDD

Domain Driven Design 领域驱动设计

DDD不是架构，而是一种架构设计方法论，通过边界划分将复杂业务领域简单化，帮我买设计出清晰的领域和应用边界。可以很容易的实现架构演进。



### 面向过程的设计方式（贫血模型）

### 面向过程的领域驱动设计方式（充血模型）

[DDD例子](https://www.cnblogs.com/qixuejia/p/10789612.html)

[VO、DTO、DO、PO的概念、区别、用处](https://www.cnblogs.com/qixuejia/p/4390086.html)



