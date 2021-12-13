package singlelinkedlist

type ICycleLinkedList interface {
	Add(value interface{}) *Element
	InsertBefore(index int, value interface{}) (*Element, error)
	Remove(index int) (*Element, error)
	Set(index int, value interface{}) (*Element, error)
	Get(index int) (*Element, error)
	ShowLinkedList() []interface{}
}

type Element struct {
	Value		interface{}
	Next 		*Element
}

func NewElement(value interface{}, next *Element) *Element {
	return &Element{
		Value: value,
		Next:  next,
	}
}

type CycleLinkedList struct {
	Size	int
	Header  *Element
}
