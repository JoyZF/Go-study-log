package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

func MyFunc(a interface{})  {
	of := reflect.TypeOf(a)
	switch of.Kind() {
	case reflect.String:
		fmt.Println(a.(string))
		a = "bbb"
		fmt.Println(a)
	case reflect.Struct:
		fmt.Println(a.(User).Name)
	}
}

func main()  {
	MyFunc("aaa")
}