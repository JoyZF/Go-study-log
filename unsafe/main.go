package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type Programmer struct {
	name string
	age int
	language string
}

func main() {
	// 错误示例
	//num := 5
	//numPointer := &num
	//
	//flnum := (*float32)(numPointer)
	//fmt.Println(flnum)
	//num := 5
	//numPointer := &num
	//flnum := (*float32)(unsafe.Pointer(numPointer))
	//fmt.Println(flnum)
	p := Programmer{
		"joy",
		18,
		"go",
	}
	fmt.Println(p)
	fmt.Println(unsafe.Pointer(&p))
	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int(1)) + unsafe.Sizeof(string(""))))
	*lang = "Golang"
	fmt.Println(p)
	sync.Map{}
}



func string2bytes(s string) []byte {
	//获取string的地址
	//获取[]byte的地址
	// 地址强转
	return *(*[]byte)(unsafe.Pointer(&s))
}
func bytes2string(b []byte) string{
	// 同理
	return *(*string)(unsafe.Pointer(&b))
}
