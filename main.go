package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	     Val int
	     Next *ListNode
}

type A interface {
	SetA()
}

type B struct {
	Age int
}


func test(x byte)  {
	fmt.Println(x)
}

const (
	x uint16 = 120
	y
	s = "abc"
	z
)
func QuickSort(slice_arg []int, iLeft int, iRight int) {
	if iLeft < iRight {
		var iTmpVal = slice_arg[iLeft]
		var i, j = iLeft, iRight
		for i < j {
			// 右边的值大于对标元素，右边的下标往左边移动 找出在右边 但小于对标元素的下标
			for i < j && slice_arg[j] > iTmpVal {
				j--
			}
			// 交换
			t := slice_arg[i]
			slice_arg[i] = slice_arg[j]
			slice_arg[j] = t
			i++
			// 左边的值小于对标元素，左边的下标往右边移动 找出在左边 但大于对标元素的下标
			for i < j && slice_arg[i] < iTmpVal {
				i++
			}
			// 交换
			t2 := slice_arg[j]
			slice_arg[j] = slice_arg[i]
			slice_arg[i] = t2
			j--
		}
		QuickSort(slice_arg, iLeft, i-1)
		QuickSort(slice_arg, j+1, iRight)
	}
}
type T struct {
	ls []int
}
func foo(t T) {
	t.ls[0] = 100
}
type People struct {
	name string `json:"name"`
}
func main() {
	m := make(map[int]string, 9)
	m[1] = "hello"
	m[2] = "world"
	m[3] = "go"
	for i := 0; i < math.MaxInt64; i++ {
		m[i] = ""
		//delete(m, i)
	}
	//v, ok := m[1]
	//_, _ = fn(v, ok)
	//delete(m, 1)
}

func fn(v string, ok bool) (string, bool) {
	return v, ok
}
func DeferTest1(i int) (r int) {
	r = i
	defer func() {
		r += 3
	}()
	return r
}

func DeferTest2(i int) (r int) {
	defer func() {
		r += i
	}()
	return 2
}

func middleNode(head *ListNode) *ListNode {
	temp := make([]*ListNode,1)
	for head.Next != nil {
		temp = append(temp,head)
		head = head.Next
	}

	return temp[len(temp) / 2+1]
}