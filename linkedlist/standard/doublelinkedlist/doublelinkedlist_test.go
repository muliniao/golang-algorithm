package doublelinkedlist

import (
	"testing"
)

var doubleLinkedList *DoubleLinkedList

func init() {
	doubleLinkedList = NewDoubleLinkedList(NewElement(1, nil, nil))
	doubleLinkedList.Add(2)
	doubleLinkedList.Add(3)
}

func TestDoubleLinkedList_Add(t *testing.T) {

}

func TestDoubleLinkedList_InsertBefore(t *testing.T) {

}

func TestDoubleLinkedList_Remove(t *testing.T) {

}

func TestDoubleLinkedList_Set(t *testing.T) {

}

func TestDoubleLinkedList_Get(t *testing.T) {

}

func TestDoubleLinkedList_ShowLinkedList(t *testing.T) {

}