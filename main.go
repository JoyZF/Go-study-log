package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func handle()  {
	path := "../test.txt"
	data, _ := ioutil.ReadFile(filepath.Join("home/user/", path))
	fmt.Println(string(data))
}



func main() {
	handle()
}