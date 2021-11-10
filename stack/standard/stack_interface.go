package standard

type Stack interface {
	Push(data interface{}) (interface{}, error)
	Pop() (interface{}, error)
	IsEmpty() bool
	Peek() (interface{},error)
 	Flush()
}
