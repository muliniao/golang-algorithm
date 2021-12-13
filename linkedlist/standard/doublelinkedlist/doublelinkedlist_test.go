package doublelinkedlist

import (
	"fmt"
	"testing"
)

func TestDoubleLinkedList_Add(t *testing.T) {

	// New
	doubleLinkedList := NewDoubleLinkedList(NewElement(1, nil, nil))

	// Add
	fmt.Println("Add is calling")
	doubleLinkedList.Add(2)
	doubleLinkedList.Add(4)
	fmt.Println(doubleLinkedList.ShowLinkedList())

	// InsertBefore
	fmt.Println("InsertBefore is calling")
	doubleLinkedList.InsertBefore(3, 3)
	fmt.Println(doubleLinkedList.ShowLinkedList())

	// Set
	fmt.Println("Set is calling")
	doubleLinkedList.Set(3,30)
	fmt.Println(doubleLinkedList.ShowLinkedList())

	// Get
	fmt.Println("Get is calling")
	fmt.Println(doubleLinkedList.Get(2))

	// Remove
	fmt.Println("Remove is calling")
	doubleLinkedList.Remove(2)
	fmt.Println(doubleLinkedList.ShowLinkedList())

	// ShowLinkedList
	fmt.Println("ShowLinkedList is calling")
	fmt.Println(doubleLinkedList.ShowLinkedList())

}
