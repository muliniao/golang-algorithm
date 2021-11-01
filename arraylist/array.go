package arraylist

import (
	"errors"
	"math"
)

const (
	DefaultCapacity = 10
	MaxArraySize    = math.MaxInt64 - 8
)

type ArrayList struct {
	elementData []interface{}
	size 		int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		elementData: make([]interface{}, 0, DefaultCapacity),
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

func (a *ArrayList) add(value interface{}) bool {
	newSize := a.size + 1
	a.ensureCapacityInternal(newSize)
	a.elementData[newSize] = value
	return true
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

func (a *ArrayList) ensureCapacityInternal(minCapacity int) {

}

func (a *ArrayList) calculateCapacity(elementData []interface{}, minCapacity int) int {
	// Java Code
	//if (elementData == DEFAULTCAPACITY_EMPTY_ELEMENTDATA) {
	//	return Math.max(DEFAULT_CAPACITY, minCapacity);
	//}
	//return minCapacity;

	return minCapacity
}

func (a *ArrayList) ensureExplicitCapacity(minCapacity int) {


}

func (a *ArrayList) grow(minCapacity int) {

	oldCapacity := len(a.elementData)
	// 加0.5倍
	newCapacity := oldCapacity + (oldCapacity >> 1)
	// ?如果加0.5倍后，还不行
	if newCapacity - minCapacity < 0 {
		newCapacity = minCapacity
	}

	if newCapacity -MaxArraySize > 0 {
		newCapacity = a.hugeCapacity(minCapacity)
	}

	copy := make([]interface{}, 0, newCapacity)
	a.elementData = a.arrayCopy(a.elementData, 0, copy, 0, newCapacity)

}

func (a *ArrayList) hugeCapacity(minCapacity int) int {
	if minCapacity < 0 {
		return 0
	}

	if minCapacity > MaxArraySize {
		return math.MaxInt64
	}
	return MaxArraySize
}

// src: source array
// srcPos: starting position in the source array
// dest: the destination array
// destPos: starting position in the destination data
// length: the number of array elements to be copied, newCapacity
func (a *ArrayList) arrayCopy(src []interface{}, srcPos int, dest []interface{}, desPos int, length int) []interface{} {
	dest = append(dest, src)
	return dest
}
