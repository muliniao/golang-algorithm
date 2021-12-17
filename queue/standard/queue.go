package standard

type IQueue interface {
	Add(object interface{}) bool
	Remove() interface{}
	Element() interface{}
}
