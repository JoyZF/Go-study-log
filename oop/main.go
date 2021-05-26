package main

import _package "study/oop/package"

func main()  {
	animal := _package.NewAnimal()
	animal.SetName("bb")
	animal.GetName()
	animal.SetName("pp")
	animal.GetName()

	cat := _package.Cat{}
	cat.GetName()
}
