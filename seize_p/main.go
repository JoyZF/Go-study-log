package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	//模拟单线程
	runtime.GOMAXPROCS(1)
	go func() {
		for  {

		}
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("joy")
	//这个例子在老版本的go中会一致阻塞
}