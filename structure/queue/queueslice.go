package queue

type QueueSlice struct {
	list   []interface{}
	length int
}

func NewQueueSlice(elems ...interface{}) *QueueSlice {
	queue := &QueueSlice{}
	for _, elem := range elems {
		queue.Enqueue(elem)
	}
	return queue
}

func (q *QueueSlice) Enqueue(elem interface{}) {
	q.list = append(q.list, elem)
	q.length++
}

func (q *QueueSlice) Dequeue() interface{} {
	if q.Length() == 0 {
		return nil
	}
	dequeued := q.list[0]
	q.list = q.list[1:]
	q.length--
	return dequeued
}

func (q *QueueSlice) Length() int {
	return q.length
}

func (q *QueueSlice) IsEmpty() bool {
	return q.Length() == 0
}
