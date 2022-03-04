package main

import "fmt"

var a = 11

func foo(n int)  {
	a := 1
	a += n
	fmt.Println(a)
}

func main() {
	fmt.Println("a = ", a)
	foo(5)
	fmt.Println("after calling foo ,a = ", a)
}
