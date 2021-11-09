package skl

import "math"

const (
	maxLevel = 18
	probability float64 = 1 / math.E
)

type  handleEle func(e * Element)  bool

type (
	Node struct {
		next []*Element
	}

	Element struct {
		Node
		key string
		value interface{}
	}

	SkipList struct {
		Node
		maxLevel int
		Len int
		randSource float64
		probability float64
		probTable []float64
		preNodesCache []*Node
	}
)

func NewSkipList() *SkipList {
	return &SkipList{
		Node:Node{next: make([]*Element,maxLevel)},
	}
}


