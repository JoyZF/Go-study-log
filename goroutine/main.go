package main

import (
	"context"
	"fmt"
	"time"
)

//func main() {
//	//type Group struct {
//	//	cancel func()
//	//
//	//	wg sync.WaitGroup
//	//
//	//	errOnce sync.Once
//	//	err     error
//	//}
//	//eg := new(errgroup.Group)
//}

//借助channel的close机制来完成对goroutine的控制
//func main() {
//	ch := make(chan string, 6)
//	go func() {
//		for true {
//			v,ok := <-ch
//			if !ok {
//				fmt.Println("结束")
//				return
//			}
//			fmt.Println(v)
//		}
//	}()
//	ch <- "123"
//	close(ch)
//	time.Sleep(time.Second)
//	ch <- "456"
//	close(ch)
//	time.Sleep(time.Second)
//}
//定期轮询 channel
//func main() {
//	ch := make(chan string, 6)
//	done := make(chan struct{})
//	go func() {
//		for  {
//			select {
//			case ch<-"123":
//			case <-done:
//				close(ch)
//				return
//			}
//			time.Sleep(100 * time.Millisecond)
//		}
//
//	}()
//	go func() {
//		time.Sleep(3 * time.Second)
//		done<- struct{}{}
//	}()
//	for i := range ch {
//		fmt.Println("接受到的值",i)
//	}
//	fmt.Println("结束")
//}
//使用context
func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for true {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("123")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	<-ch
	fmt.Println("结束")
}