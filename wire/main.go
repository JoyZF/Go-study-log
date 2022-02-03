package wire

import "fmt"

type Foo struct {

}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
	foo *Foo
}

func NewBar(foo *Foo) *Bar {
	return &Bar{
		foo: foo,
	}
}

func (b *Bar) Test() {
	fmt.Println("hello")
}


