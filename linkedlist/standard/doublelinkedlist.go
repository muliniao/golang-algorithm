package standard

import (
	"errors"
)

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

type DoubleLinkedList struct {
	Size	int
	Header  *Element
}

func NewDoubleLinkedList(header *Element) *DoubleLinkedList {
	return &DoubleLinkedList{
		Size:   1,
		Header: header,
	}
}

func (d *DoubleLinkedList) Add(value interface{}) *Element {

	if d.Size == 0 {
		d.Header = NewElement(value, nil, nil)
		d.Size ++
		return d.Header
	}

	temp := d.Header
	for i := 1; i < d.Size; i++ {
		temp  = temp.Next
	}

	temp.Next = NewElement(value, nil, temp)
	d.Size ++

	return temp.Next
}

func (d *DoubleLinkedList) InsertBefore(index int, value interface{}) (*Element, error) {

	err := d.checkElementIndex(index)
	if err != nil {
		return nil, err
	}

	temp := d.Header
	for i := 1; i < index; i++ {
		temp = temp.Next
	}

	newElement := NewElement(value, nil, nil)
	newElement.Next = temp
	newElement.Prev = temp.Prev
	temp.Prev.Next = newElement
	temp.Prev = newElement

	d.Size ++
	return temp, nil
}

func (d *DoubleLinkedList) Remove(index int) (*Element, error) {

	err := d.checkElementIndex(index)
	if err != nil {
		return nil, err
	}

	temp := d.Header
	for i := 1; i < index; i++ {
		temp = temp.Next
	}

	temp.Prev.Next = temp.Next
	temp.Next.Prev = temp.Prev
	d.Size--
	return temp, err
}

func (d *DoubleLinkedList) Set(index int, value interface{}) (*Element, error) {

	err := d.checkElementIndex(index)
	if err != nil {
		return nil, err
	}

	temp := d.Header
	for i := 1; i < index; i++ {
		temp = temp.Next
	}

	temp.Value = value
	return temp, err
}

func (d *DoubleLinkedList) Get(index int) (*Element, error) {

	err := d.checkElementIndex(index)
	if err != nil {
		return nil, err
	}

	temp := d.Header
	for i := 1; i < index; i++ {
		temp = temp.Next
	}

	return temp, err
}

func (d *DoubleLinkedList) ShowLinkedList() []interface{} {

	list := make([]interface{}, 0)

	for temp := d.Header; temp != nil; temp = temp.Next {
		list = append(list, temp.Value)
	}
	return list
}

func (d *DoubleLinkedList) checkElementIndex(index int) error {
	if index < 0 || index > d.Size {
		return errors.New("index is out of size")
	}
	return nil
}
