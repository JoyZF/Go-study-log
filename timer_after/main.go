package main

import (
	"fmt"
	"time"
)

//func main()  {
//	ch := make(chan string)
//	go func() {
//		time.Sleep(time.Second * 3)
//		ch <- "joy"
//	}()
//
//	select {
//	case _ = <-ch:
//	case <-time.After(time.Second * 1):
//		fmt.Println("bye bye time out")
//	}
//}

//func main()  {
//	ch := make(chan int, 10)
//	go func() {
//		in := 1
//		for  {
//			in++
//			ch <- in
//		}
//	}()
//
//	for  {
//		select {
//		case _ = <-ch:
//			continue
//		case <-time.After(3 * time.Minute):
//			//每次进行select时都会重新初始化一个全新的计时器
//			//这个计时器时在3分钟之后才会被触发指向，
//			//但是计时器激活后又没有引用关系 所以被GC掉了
//			//但是time.After的定时任务还在时间堆中等待触发，所以它永远不会到期
//			//成为孤以致泄露
//			fmt.Printf("现在是：%d",time.Now().Unix())
//		}
//	}
//}

//正确写法
func main()  {
	timer := time.NewTimer(3 * time.Minute)
	defer  timer.Stop()
	ch := make(chan int, 10)
	go func() {
		in := 1
		for  {
			in++
			ch <- in
		}
	}()
	for  {
		select {
		case _ = <-ch:
		case <-timer.C:
			fmt.Printf("现在时 %d",time.Now().Unix())
		}
	}
}