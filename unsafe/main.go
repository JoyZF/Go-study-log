package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 错误示例
	//num := 5
	//numPointer := &num
	//
	//flnum := (*float32)(numPointer)
	//fmt.Println(flnum)
	num := 5
	numPointer := &num
	flnum := (*float32)(unsafe.Pointer(numPointer))
	fmt.Println(flnum)
}