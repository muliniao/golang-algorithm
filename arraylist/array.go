package arraylist

import "errors"

const DefaultCapacity = 10

type ArrayList struct {
	elementData []interface{}
	size 		int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		elementData: nil,
		size:        0,
	}
}

func (a *ArrayList) get(index int) interface{} {
	//rangeCheck(index)
	return a.elementData[index]
}

func (a *ArrayList) set() {

}

func (a *ArrayList) add() {

}

func (a *ArrayList) remove() {

}

func (a *ArrayList) indexOf() {

}

func (a *ArrayList) lastIndexOf() {

}

func (a *ArrayList) clear() {

}

func (a *ArrayList) rangeCheck(index int) error {
	if index > a.size {
		return errors.New("index out of bounds")
	}
	return nil
}
