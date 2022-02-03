package main

import (
	"fmt"
)

type ListNode struct {
	     Val int
	     Next *ListNode
}
func main() {
	name := "hello world"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c", name[i])
	}
}


func middleNode(head *ListNode) *ListNode {
	temp := make([]*ListNode,1)
	for head.Next != nil {
		temp = append(temp,head)
		head = head.Next
	}

	return temp[len(temp) / 2+1]
}