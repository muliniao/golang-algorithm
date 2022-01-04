package main

import (
	"container/list"
	"fmt"
)

func main() {

	linkedList := list.New()

	linkedList.PushFront("aaaaaaa")
	linkedList.PushFront("bbbbbbb")
	linkedList.PushFront("ccccccc")
	linkedList.PushFront("ddddddd")
	linkedList.PushFront("eeeeeee")

	fmt.Println(linkedList.Len())
	fmt.Println(linkedList)

}
