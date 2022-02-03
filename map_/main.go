package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	m := make(map[int32]string)
	m[0] = "EDDYCJY1"
	m[1] = "EDDYCJY2"
	m[2] = "EDDYCJY3"
	m[3] = "EDDYCJY4"
	m[4] = "EDDYCJY5"

	for k, v := range m {
		log.Printf("k: %v, v: %v", k, v)
	}
	strings := make(chan string)
	go func() {
		// 仅能用于channel发送和接收消息的专用语句，运行期间是阻塞的 所以也可以用来阻塞监听goroutine
		// select 在语言层面上提供了多路复用
		select {
		case <-strings:
			s := <-strings
			fmt.Println(s)
		default:

		}

	}()
	for i := 0; i < 100000; i++ {
		strings <- time.Now().String()
	}

	//之所以map不是有序的 是因为每次重新for range map 的起始位置根本就不固定！
}