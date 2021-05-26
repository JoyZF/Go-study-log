package main

import (
	"fmt"
	"net/rpc"
)

func main()  {
	client, err := rpc.DialHTTP("tcp", "localhost:8082")
	if err != nil {
		panic(err.Error())
	}

	var req float32 //请求值
	req = 3

	var resp *float32 //返回值
	//上述的调用方法核心在于client.Call方法的调用，该方法有三个参数，第一个参数表示要调用的远端服务的方法名，第二个参数是调用时要传入的参数，第三个参数是调用要接收的返回值。 上述的Call方法调用实现的方式是同步的调用，
	err = client.Call("MathUtil.CalculateCircleArea", req, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)


	//var respSync *float32
	////异步的调用方式
	//syncCall := client.Go("MathUtil.CalculateCircleArea", req, &respSync, nil)
	//replayDone := <-syncCall.Done
	//fmt.Println(replayDone)
	//fmt.Println(*respSync)

}