package arraylist

type Array interface {
	Get(index int) (interface{}, error)
	Set(index int, value interface{}) (interface{}, error)
	Add(value interface{}) bool
	Remove(index int) (interface{}, error)
	Clear()
}
