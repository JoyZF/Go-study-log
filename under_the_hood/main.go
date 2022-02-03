package main
//一个包内的 init 函数的调用顺序取决于声明的顺序，即从上而下依次调用。
import (
	"fmt"
	_ "study/under_the_hood/init2"
	"time"
)
import _ "study/under_the_hood/init1"

func init()  {

}

func main()  {
	write()
}

func write()  {
	var step int64 = 1000000
	var t1 time.Time
	m := map[int64]int64{}
	for i := int64(0); ; i+=step {
		t1 = time.Now()
		for j := int64(0); j < step; j++ {
			m[i+j] = i + j
		}
		fmt.Printf("%d done, time: %v\n", i, time.Since(t1).Seconds())
	}
}
