package main

import (
	"fmt"
	"time"
)

//func main() {
//	//为什么要使用空结构体
//	//ch := make(chan struct{})
//	//为了节省空间
//	//Go语言中宽度描述了一个类型的实例所占用的存储空间的字节数
//	//宽度是一个类型的属性，在go语言中每个值都有一个类型，值的宽度由其类型定义，并且总是8bits的倍数
//	var a int
//	var b string
//	var c bool
//	var d [3]int32
//	var e []string
//	var f map[string]bool
//	fmt.Println(
//		unsafe.Sizeof(a),
//		unsafe.Sizeof(b),
//		unsafe.Sizeof(c),
//		unsafe.Sizeof(d),
//		unsafe.Sizeof(e),
//		unsafe.Sizeof(f),
//		)
//	//单单声明变量就用掉了8 16 1 12 24 8 的长度
//	var g struct{}
//	fmt.Println(unsafe.Sizeof(g))
//	//而用空结构体的话 长度位 0
//	//完美切合人们对占位符的基本诉求，就是占着坑位，满足基本输入输出就好。
//	//为什么空结构体的宽度为0？
//	// base address for all 0-byte allocations
//	//var zerobase uintptr
//	//
//	//func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
//	//	...
//	//	if size == 0 { Go编译器在内存分配时做了优化 当发现size为0时，会直接返回变量zerobase的引用 ，该变量是所有 0 字节的基准地址，不占据任何宽度。
//	//	return unsafe.Pointer(&zerobase)
//	//}
//	//}
//	//因此空结构的广泛使用时开发者借助了这个小优化，达到占位符的目的
//	//使用场景
//	//- 实现方法接受者
//	//- 实现集合类型
//	//- 实现空通道
//
//
//
//}

//实现方法接收者 优势 易扩展、省空间、最结构化
//type T struct {
//
//}
//
//func (t *T) Call()  {
//	fmt.Println("空结构体实现方法接收")
//}
//
//func main()  {
//	var t T
//	t.Call()
//}


//实现集合类型
//在Go语言中没有提供set集合的实现，因为我们可以使用map代替
//但是map需要用到key 和value 实际上set不需要用到value  这个时候可以用空结构体代替value
//type Set map[string]struct{}
//
//func (s Set) Append(k string)  {
//	s[k] = struct{}{}
//}
//
//func (s Set) Remove(k string)  {
//	delete(s,k)
//}
//
//func (s Set) Exist(k string) bool  {
//	if _,ok := s[k];ok {
//		return true
//	}else{
//		return false
//	}
//}
//
//func main() {
//	s := Set{}
//	s.Append("空结构体")
//	s.Append("空结构头")
//	s.Append("空街头")
//	s.Remove("空街头")
//	fmt.Println(s.Exist("空街头"))
//}

//实现空通道
//Go channel 的使用场景中，常常会遇到通知型channel,其不需要发送任何数据，只是用于协调Goroutine的运行，用于刘庄各类状态或是控制并发情况
func main() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	fmt.Println("空结构体用于")
	<-ch
	fmt.Println("空通道")
}
//由于该 channel 使用的是空结构体，因此也不会带来额外的内存开销。
