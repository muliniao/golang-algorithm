package main

import (
	"container/list"
	"fmt"
)

func main() {

	var lists []*list.List

	lists = make([]*list.List, 0)

	linkedList001 := list.New()
	linkedList001.PushFront("aaaaaa")
	linkedList001.PushFront("bbbbbb")
	linkedList001.PushFront("cccccc")
	linkedList001.PushFront("dddddd")

	linkedList002 := list.New()
	linkedList002.PushFront("111111")
	linkedList002.PushFront("222222")
	linkedList002.PushFront("333333")

	lists = append(lists, linkedList001, linkedList002)

	fmt.Println(lists[0])
	fmt.Println(lists[1])

}
