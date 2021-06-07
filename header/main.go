package main


import (
	"fmt"
	"reflect"
	"unsafe"
)

// SliceHeader StringHeader

//SliceHeader 是Go slice 的运行时表现
//type SliceHeader struct {
//	Data uintptr
//	Len  int
//	Cap  int
//}

//type StringHeader struct {
//	Data uintptr
//	Len  int
//}
func main() {
	s := [4]string{"1","2","3","4"}
	s1 := s[0:1]
	s2 := s[:]

	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	sh2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	fmt.Println(sh1.Len, sh1.Cap, sh1.Data)
	fmt.Println(sh2.Len, sh2.Cap, sh2.Data)
	//这其实是 Go 语言本身为了减少内存占用，提高整体的性能才这么设计的。
	//将切片复制到任意函数的时候，对底层数组大小都不会影响。复制时只会复制切片本身（值传递），不会涉及底层数组。
	//也就是在函数间传递切片，其只拷贝 24 个字节（指针字段 8 个字节，长度和容量分别需要 8 个字节），效率很高。
	//假设在没有超过容量（cap）的情况下，对第二个切片操作会影响第一个切片。

	//StringHeader
	//字符串运行时表现
	ss := "1"
	ss1 := "1"
	ss2 := "2"[1:]
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&ss)).Data)
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&ss1)).Data)
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&ss2)).Data)
	//在Go中字符串都是只读的 为了节省内存，相同字面量的字符串通常对应于同一字符串常量，因此指向同一个底层数组。

	//SliceHeader StringHeader 可以利用其实现零拷贝的string到bytes的转换
}

//错误例子 因此在上述代码中会出现将 Data 作为值拷贝的情况，这就会导致无法保证它所引用的数据不会被垃圾回收（GC）。
func string2bytes(s string) []byte  {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func string2bytes1(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	var b []byte
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pbytes.Data = stringHeader.Data
	pbytes.Len = stringHeader.Len
	pbytes.Cap = stringHeader.Len

	return b
}
//在性能方面，若只是期望单纯的转换，对容量（cap）等字段值不敏感，也可以使用以下方式：
func string2bytes2(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}