package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)
var datas []string

func init()  {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
}

func main()  {
	//var m sync.Mutex
	//var datas = make(map[int]struct{})
	//for i:=0;i<999;i++ {
	//	go func(i int) {
	//		m.Lock()
	//		defer  m.Unlock()
	//		datas[i] = struct{}{}
	//	}(i)
	//}
	//
	//_ = http.ListenAndServe(":6061", nil)
	go func() {
		for  {
			log.Printf("len : %d",Add("go-grogramming-tour-book"))
			time.Sleep(time.Second * 10)
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}