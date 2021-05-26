package main

import "fmt"

type People struct {

}

type MyStruct struct {
	Name string
}

func (s MyStruct) SetName1(name string)  {
	s.Name = name
}
func (s *MyStruct) SetName2(name string)  {
	s.Name = name
}

func main()  {
	//内存分配在栈 在Go编译器的代码优化阶段，会对其进行优化，直接返回false 并不是真正去比较了 可以通过 go run -gcflags="-N -l" main.go 不让他优化
	a := new(struct{})
	b := new(struct{})
	println(a, b, a == b)
	//fmt.Println 之后用到了引用 导致内存逃逸到堆上，空结构体默认分配的是runtime.zerobase 变量，是专门用于分配到堆上的 0 字节基础地址。因此两个空结构体，都是 runtime.zerobase，一比较当然就是 true 了。
	c := new(struct{})
	d := new(struct{})
	fmt.Println(c, d)
	println(c, d, c == d)
	//如果选中指针参数还是值参数
	/**
	 * 1、在使用上的考虑：方法是否需要修改接收器？如果需要，接收器必须是一个指针
	 * 2、在效率上的考虑：如果接收器是个大的结构，使用指针接收器会好很多
	 * 3、在一致性上的考虑：如果类型的某些方法必须有指针接收器，那么其余的方法也应该有指针接收器，所以无论类型如何使用，方法集都是一致的。
	 * 因此除非方法的语义需要指针，那么值接收器是最高效和清晰的。在 GC 方面，也不需要过度关注。出现时再解决就好了。
	*/

	var myStruct MyStruct
	myStruct.SetName1("joy")
	fmt.Println(myStruct.Name)
	myStruct.SetName1("zhangsan")
	fmt.Println(myStruct.Name)
}
