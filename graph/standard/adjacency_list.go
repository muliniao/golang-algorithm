package standard

import "container/list"

type AdjacencyList struct {
	list []*list.List
	size int
}

func NewAdjacencyList(size int) *AdjacencyList {
	return &AdjacencyList{
		list: make([]*list.List, 0, size),
		size: size,
	}
}
