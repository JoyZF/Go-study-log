package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct{
}
//该方法向外暴露：提供计算圆形面积的服务
func (mu *MathUtil) CalculateCircleArea(req float32, resp *float32) error {
	*resp = math.Pi * req * req //圆形的面积 s = π * r * r
	return nil //返回类型
}


//经过服务注册和监听处理，RPC调用过程中的服务端实现就已经完成了。接下来需要实现的是客户端请求代码的实现。
func main()  {
	mathUtil := new(MathUtil)
	err := rpc.Register(mathUtil)
	if err != nil {
		panic(err.Error())
	}

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err.Error())
	}
	go http.Serve(listen,nil)
}