package standard

import (
	"fmt"
	"testing"
)

func TestDoubleLinkedList_Add(t *testing.T) {

	// New
	doubleLinkedList := NewDoubleLinkedList(NewElement(1, nil, nil))

	// Add
	doubleLinkedList.Add(2)
	doubleLinkedList.Add(4)
	fmt.Println(doubleLinkedList.ShowLinkedList())

	// InsertBefore
	doubleLinkedList.InsertBefore(3, 3)
	fmt.Println(doubleLinkedList.ShowLinkedList())

	// Set
	doubleLinkedList.Set(3,30)
	fmt.Println(doubleLinkedList.ShowLinkedList())

	// Get
	fmt.Println(doubleLinkedList.Get(2))

	// Remove
	doubleLinkedList.Remove(2)
	fmt.Println(doubleLinkedList.ShowLinkedList())

	// ShowLinkedList
	fmt.Println(doubleLinkedList.ShowLinkedList())

}
