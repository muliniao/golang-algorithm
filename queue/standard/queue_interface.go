package standard

type Queue interface {
	Add(object interface{}) bool
	Remove() interface{}
	Element() interface{}
}
