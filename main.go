package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

func StringToByte(key *string) []byte {
	strPtr := (*reflect.SliceHeader)(unsafe.Pointer(key))
	strPtr.Cap = strPtr.Len
	b := *(*[]byte)(unsafe.Pointer(strPtr))
	return b
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}