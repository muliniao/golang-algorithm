package arraystack

import (
	"fmt"
	"testing"
)

var stack *ArrayStack

func init() {
	stack = NewArrayStack()
	stack.Push("111")
	stack.Push("222")
	stack.Push("333")
	stack.Push("444")
	stack.Push("555")
}

func TestArrayStack_Push(t *testing.T) {
	stack.Push("666")
	stack.Push("777")
	fmt.Println(stack)
}

func TestArrayStack_Peek(t *testing.T) {
	data, err := stack.Peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

func TestArrayStack_Pop(t *testing.T) {
	data, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stack)
	fmt.Println(data)
}

func TestArrayStack_IsEmpty(t *testing.T) {
	fmt.Println(stack.IsEmpty())
}

func TestArrayStack_Search(t *testing.T) {
	index, err := stack.Search("333")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(index)
}
