package main

import "fmt"

func main()  {
	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		fmt.Println(i, c)
	}
}

func fori()  {
	var times = [5]int{1,2,3,4,5}
	for range times {

	}
}

func forrange()  {
	var times = [5]int{1,2,3,4,5}
	for _ = range times {

	}
}
