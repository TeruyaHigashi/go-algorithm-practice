package queue

type QueueLinkedList struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	value interface{}
	next  *node
}

func NewQueueLinkedList(elems ...interface{}) *QueueLinkedList {
	queue := &QueueLinkedList{}
	if elems == nil {
		return queue
	}
	for _, elem := range elems {
		queue.Enqueue(elem)
	}
	return queue
}

func (q *QueueLinkedList) Enqueue(elem interface{}) {
	enqueued := &node{value: elem, next: nil}
	if q.length == 0 {
		q.head = enqueued
	} else {
		q.tail.next = enqueued
	}
	q.tail = enqueued
	q.length++
}

func (q *QueueLinkedList) Dequeue() interface{} {
	if q.Length() == 0 {
		return nil
	}
	dequeued := q.head.value
	if q.Length() == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.length--
	return dequeued
}

func (q *QueueLinkedList) Length() int {
	return q.length
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.Length() == 0
}
