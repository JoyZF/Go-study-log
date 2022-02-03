package main

import (
	"fmt"
)

type Animal interface {
	eat()
	run()
}

type Cat struct {
	Name string
	Sex string
}

type Dog struct {
	Name string
}

func (receiver Cat) eat()  {
	fmt.Println(receiver.Name + " is eating")
}

func (receiver Cat) run()  {
	fmt.Println(receiver.Name + " is running")
}

func (receiver Dog) eat() {
	fmt.Println(receiver.Name + " is eating")
}

func (receiver Dog) run() {
	fmt.Println(receiver.Name + " is running")
}

var L Animal


type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (receiver Gopher) code()  {
	fmt.Println("coding")
}

func (receiver *Gopher) debug()  {
	fmt.Println("debug")
}

type myWriter struct {

}

//func (w myWriter) Write(p []byte) (n int, err error) {
//	return
//}

func main()  {
	//cat := Cat{
	//	Name: "Tom",
	//	Sex: "man",
	//}
	//var a Animal
	//a = cat
	//a.run()
	//dog := Dog{
	//	Name: "Spark",
	//}
	//// 多态
	//MyFunc(cat)
	//MyFunc(dog)
	//// 解耦
	//L = cat
	//L.run()
	//var c coder = &Gopher{"go"}
	//c.code()
	//c.debug()
	//var _ io.Writer = (*myWriter)(nil)
	//var _ io.Writer = myWriter{}
	var a interface{}
	var b string = "123"
	a = b
	val , ok := a.(string)
	if ok {
		fmt.Println(val)
	}
}

func MyFunc(a Animal)  {
	a.run()
	a.eat()
}
