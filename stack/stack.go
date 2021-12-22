package stack

type IStack interface {
	Push(data interface{}) (interface{}, error)
	Pop() (interface{}, error)
	IsEmpty() bool
	Peek() (interface{}, error)
	Search(data interface{}) (int, error)
}
