package main

import (
	"container/list"
	"fmt"
)

func main() {

	var lists []*list.List

	lists = make([]*list.List, 0)

	linkedList001 := list.New()
	linkedList001.PushBack("aaaaaa")
	linkedList001.PushBack("bbbbbb")
	linkedList001.PushBack("cccccc")
	linkedList001.PushBack("dddddd")

	linkedList002 := list.New()
	linkedList002.PushFront("111111")
	linkedList002.PushFront("222222")
	linkedList002.PushFront("333333")

	lists = append(lists, linkedList001, linkedList002)

	fmt.Println(lists[0])
	fmt.Println(lists[1])

}
