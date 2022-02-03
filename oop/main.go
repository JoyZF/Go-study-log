package main

import _package "study/oop/package"

type Base struct {
	Name string `json:"name"`
}

type Custom struct {
	Base
	Sex string `json:"sex"`
}

func Test(base Base)  {

}

func main()  {
	animal := _package.NewAnimal()
	animal.SetName("bb")
	animal.GetName()
	animal.SetName("pp")
	animal.GetName()

	cat := _package.Cat{}
	cat.GetName()
	custom := Custom{}
}


