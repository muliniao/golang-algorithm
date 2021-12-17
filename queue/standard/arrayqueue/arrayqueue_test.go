package arrayqueue

import (
	"fmt"
	"testing"
)

var arrayQueue *ArrayQueue

func init() {
	arrayQueue = NewArrayQueue()
	arrayQueue.Add("111")
	arrayQueue.Add("222")
}

func TestArrayQueue_Add(t *testing.T) {
	arrayQueue.Add("333")
	fmt.Println(arrayQueue)
}

func TestArrayQueue_Remove(t *testing.T) {
	removedElement := arrayQueue.Remove()
	fmt.Println(arrayQueue)
	fmt.Println(removedElement)
}

func TestArrayQueue_Element(t *testing.T) {
	headElement := arrayQueue.Element()
	fmt.Println(headElement)
}
