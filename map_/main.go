package main

import "log"

func main() {
	m := make(map[int32]string)
	m[0] = "EDDYCJY1"
	m[1] = "EDDYCJY2"
	m[2] = "EDDYCJY3"
	m[3] = "EDDYCJY4"
	m[4] = "EDDYCJY5"

	for k, v := range m {
		log.Printf("k: %v, v: %v", k, v)
	}
	//之所以map不是有序的 是因为每次重新for range map 的起始位置根本就不固定！
}