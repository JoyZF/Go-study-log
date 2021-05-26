package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

//1.1
//func main()  {
//	var name string
//	flag.StringVar(&name , "name","Go语言编程之旅","帮助信息")
//	flag.StringVar(&name,"n","Go语言编程之旅","帮助信息")
//	flag.Parse()
//	log.Printf("name : %s",name)
//}

//1.2
//var name string
//
//func main()  {
//	flag.Parsed()
//	args := flag.Args()
//	if len(args) <= 0 {
//		return
//	}
//
//	switch args[0] {
//	case "go":
//		gocmd := flag.NewFlagSet("go", flag.ExitOnError)
//		gocmd.StringVar(&name,"name","GO语言","帮助信息")
//		_ = gocmd.Parse(args[1:])
//	case "php":
//		gocmd := flag.NewFlagSet("php", flag.ExitOnError)
//		gocmd.StringVar(&name,"name","PHP语言","帮助信息")
//		_ = gocmd.Parse(args[1:])
//	}
//	log.Printf("name: %s",name)
//
//}

//1.3

type Name string
func (i *Name) String() string  {
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}
	*i = Name("joy:"+value)
	return nil
}

func main()  {
	var name Name
	flag.Var(&name,"name","帮助信息")
	flag.Parse()
	log.Printf("name : %s",name)
}