package queue

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Length() int
	IsEmpty() bool
}
