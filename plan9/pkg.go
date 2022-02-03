package main

import (
	"fmt"
	"runtime"
)

func main() {
	var buf = make([]byte,64)
	var stk = buf[:runtime.Stack(buf, false)]
	fmt.Println(string(stk))
}