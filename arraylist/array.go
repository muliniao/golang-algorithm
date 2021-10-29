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

func (a *ArrayList) get(index int) (interface{}, error) {
	if err := a.rangeCheck(index); err != nil {
		return nil, err
	}

	return a.elementData[index], nil
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

// src: source array
// srcPos: starting position in the source array
// dest: the destination array
// destPos: starting position in the destination data
// length: the number of array elements to be copied, newCapacity

func (a *ArrayList) arrayCopy(src interface{}, srcPos int, dest interface{}, desPos int, length int) {

	a.elementData = make([]interface{}, length, length)
	a.elementData = append(a.elementData, src)


}
