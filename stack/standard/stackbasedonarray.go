package standard

import "errors"

const DefaultCapacity = 128

type ArrayStack struct {
	elementData		[]interface{}
	elementCount	int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		elementData:  make([]interface{}, 0, DefaultCapacity),
		elementCount: 0,
	}
}

func (a *ArrayStack) Push(data interface{}) (interface{}, error) {
	if a.size() == DefaultCapacity{
		return nil, errors.New("stack is full")
	}

	a.addElement(data)
	return data, nil
}

func (a *ArrayStack) Pop() (interface{}, error) {
	len := a.size()
	object, err := a.Peek()
	if err != nil {
		return nil, err
	}

	a.removeElementAt(len - 1)
	return object, nil
}

func (a *ArrayStack) IsEmpty() bool {
	return a.size() == 0
}

func (a *ArrayStack) Peek() (interface{}, error) {
	len := a.size()
	if a.IsEmpty() {
		return nil, errors.New("empty stack exception")
	}

	return a.elementAt(len - 1)
}

func (a *ArrayStack) Search(data interface{}) (int, error) {
	i, err := a.lastIndexOf(data, a.elementCount - 1)
	if err != nil {
		return -1, err
	}

	if i >= 0 {
		return a.size() - i, nil
	}

	return -1, nil
}

func (a *ArrayStack) addElement(data interface{}) {
	a.elementCount++
	a.elementData = append(a.elementData, data)
}

func (a *ArrayStack) elementAt(index int) (interface{}, error) {
	if index >= a.elementCount {
		return nil, errors.New("array index out of bounds")
	}

	return a.elementData[index], nil
}

func (a *ArrayStack) removeElementAt(index int) error {
	if index >= a.elementCount || index < 0 {
		return errors.New("array index out of bounds")
	}

	a.elementCount--
	a.elementData = append(a.elementData[:index], a.elementData[index + 1:]...)
	return nil
}

func (a *ArrayStack) lastIndexOf(data interface{}, index int) (int, error) {
	if index > a.elementCount {
		return -1, errors.New("index out of bound exception")
	}

	if data == nil {
		for i := index; i >= 0; i-- {
			if a.elementData[i] == nil {
				return i, nil
			}
		}
	} else {
		for i := index; i >= 0;	i-- {
			if a.elementData[i] == data {
				return i, nil
			}
		}
	}

	return -1, nil
}

func (a *ArrayStack) size() int {
	return a.elementCount
}
