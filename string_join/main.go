package main

import "fmt"

func main()  {
	fmt.Println(mul(1,2))
	//var a = "hello "
	//var b = "world"
	//var start = time.Now()
	//for i := 0; i < math.MaxInt16; i++ {
	//	//244000111
	//	//a = a + b
	//	//248209144
	//	//a += b
	//	//491741
	//	//builder.WriteString(b)
	//	//122207
	//	//a = strings.Join([]string{a},b)
	//	//510255788
	//	a = fmt.Sprintf("%s%s",a,b)
	//}
	//var end = time.Now()
	//sub := end.Sub(start)
	//println(sub)
}

func mul(a, b int) (c int) {
	c = a + b
	return
}