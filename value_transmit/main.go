package main

import "fmt"

//func main()  {
//	//Go语言是值传递，向一个函数传递一个int值，会得到int的副本
//	//而传递一个指针值就会得到指针的副本，但不会得到它所指向的数据
//	s := "joy"
//	fmt.Printf("main 内存地址: %p \n",&s)
//	hello(&s)
//}
//
//func hello(s *string)  {
//	fmt.Printf("hello 内存地址: %p \n",&s)
//}

//func main()  {
//	s := "joy"
//	fmt.Printf("main 内存地址： %p \n",&s)
//	hello(&s)
//	fmt.Println(s)
//}
//
//func hello(s *string)  {
//	fmt.Printf("hello 内存地址: %p \n",&s)
//	*s = "hello joy"
//}

func main()  {
	m := make(map[string]string)
	m["joy"] = "hello"
	fmt.Printf("main 内存地址: %p \n",&m)
	hello(m)
	fmt.Printf("%v",m)
}

func hello(p map[string]string)  {
	fmt.Printf("hello 内存地址:%p \n",&p)
	p["joy"] = "bye bye joy"
}