package cyclequeue

const MinInitialCapacity = 8

// 队满: (tail + 1) % n == head
// 队空: head == tail

type CycleQueue struct {
	Elements []interface{}
	Head     int
	Tail     int
}

func NewCycleQueue() *CycleQueue {
	return &CycleQueue{
		Elements: make([]interface{}, 0, MinInitialCapacity),
		Head:     0,
		Tail:     0,
	}
}

func (c *CycleQueue) Add(object interface{}) bool {
	if (c.Tail+1)%MinInitialCapacity == c.Head {
		return false
	}

	c.Elements[c.Tail] = object
	c.Tail = (c.Tail + 1) % MinInitialCapacity
	return true
}

func (c *CycleQueue) Remove() interface{} {

	headElement := c.Element()
	c.Head = (c.Head + 1) % MinInitialCapacity
	return headElement
}

func (c *CycleQueue) Element() interface{} {
	if c.Head == c.Tail {
		return nil
	}

	return c.Elements[c.Head]
}
