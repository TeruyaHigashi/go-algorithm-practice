package linkedlist

type LinkedList interface {
	AddBegin(elem interface{})
	AddEnd(elem interface{})
	RemoveBegin() interface{}
	RemoveEnd() interface{}
	Length() int
	IsEmpty() bool
}
