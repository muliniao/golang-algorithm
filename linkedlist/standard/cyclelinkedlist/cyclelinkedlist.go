package cyclelinkedlist

import "errors"

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
	Next, Prev 	*Element
}

func NewElement(value interface{}, next, before *Element) *Element {
	return &Element{
		Value: value,
		Next:  next,
		Prev:  before,
	}
}

type CycleLinkedList struct {
	Size	int
	Header  *Element
}

func NewCycleLinkedList(header *Element) (*CycleLinkedList, error) {
	if header.Next != header.Prev {
		return nil, errors.New("it is not cycle linked list")
	}

	return &CycleLinkedList{
		Size:   1,
		Header: header,
	}, nil
}

func (c *CycleLinkedList) Add(value interface{}) *Element {
	if c.Size == 0 {
		c.Header = NewElement(value, c.Header, c.Header)
		c.Size ++
		return c.Header
	}

	temp := c.Header
	for i := 1; i < c.Size; i++ {
		temp  = temp.Next
	}

	temp.Next = NewElement(value, c.Header, temp)
	c.Header.Prev = temp.Next

	c.Size ++
	return temp.Next
}

func (c *CycleLinkedList) InsertBefore(index int, value interface{}) (*Element, error) {
	if err := c.checkElementIndex(index); err != nil {
		return nil, err
	}

	temp := c.Header
	for i := 1; i < index; i++ {
		temp = temp.Next
	}

	newElement := NewElement(value, nil, nil)
	newElement.Next = temp
	newElement.Prev = temp.Prev
	temp.Prev = newElement
	temp.Prev.Next = newElement
	c.Size ++

	return newElement, nil
}

func (c *CycleLinkedList) Remove(index int) (*Element, error) {
	if err := c.checkElementIndex(index); err != nil {
		return nil, err
	}

	temp := c.Header
	for i := 1; i < index; i++ {
		temp = temp.Next
	}

	temp.Prev.Next = temp.Next
	temp.Next.Prev = temp.Prev
	c.Size --

	return temp, nil
}

func (c *CycleLinkedList) Set(index int, value interface{}) (*Element, error) {
	if err := c.checkElementIndex(index); err != nil {
		return nil, err
	}

	temp := c.Header
	for i := 1; i < index; i++ {
		temp = temp.Next
	}

	temp.Value = value
	return temp, nil
}

func (c *CycleLinkedList) Get(index int) (*Element, error) {
	if err := c.checkElementIndex(index); err != nil {
		return nil, err
	}

	temp := c.Header
	for i := 1; i < index; i++ {
		temp = temp.Next
	}

	return temp, nil
}

func (c *CycleLinkedList) ShowLinkedList() []interface{} {

	list := make([]interface{}, 0)

	for temp := c.Header; temp != nil; temp = temp.Next {
		list = append(list, temp.Value)
	}

	return list
}

func (c *CycleLinkedList) checkElementIndex(index int) error {
	if index < 0 || index > c.Size {
		return errors.New("index is out of size")
	}
	return nil
}





