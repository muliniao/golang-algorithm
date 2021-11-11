package standard

const MinInitialCapacity = 8

type ArrayQueue struct {
	Elements	[]interface{}
	Head		int
	Tail		int
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		Elements: make([]interface{}, 0, MinInitialCapacity),
		Head:     0,
		Tail:     0,
	}
}

func (a *ArrayQueue) Add(object interface{}) bool {

	if a.Tail == MinInitialCapacity && a.Head == 0 {
		return false
	}

	// 数据搬移
	for i := a.Head; i < a.Tail; i++ {
		a.Elements[i - a.Head]	= a.Elements[i]
	}

	a.Tail -= a.Head
	a.Head = 0

	// 插入数据
	a.Elements[a.Tail] = object
	a.Tail++
	return true
}

func (a *ArrayQueue) Remove() interface{} {

	headElement := a.Element()
	a.Elements[a.Head] = nil
	a.Head++
	return headElement
	
}

func (a *ArrayQueue) Element() interface{} {

	if a.Head == a.Tail {
		return nil
	}

	return a.Elements[a.Head]
}



