package main

import "fmt"

const c = 123


var d = 456

func main()  {
	a := 1
	a++
	b := 2
	b = a
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}


