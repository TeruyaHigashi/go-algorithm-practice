package queue

import "testing"

func TestQueueSlice(t *testing.T) {

	sut := &QueueSlice{}
	lengthTest(sut, 0, t)

	sut = &QueueSlice{}
	isEmptyTest(sut, true, t)

	sut = &QueueSlice{}
	enqueueTest(sut, t)

	sut = &QueueSlice{}
	dequeueTest(sut, nil, t)

	sut = NewQueueSlice("aaa", "bbb", 123)
	lengthTest(sut, 3, t)

	sut = NewQueueSlice("aaa", "bbb", 123)
	isEmptyTest(sut, false, t)

	sut = NewQueueSlice("aaa", "bbb", 123)
	dequeueTest(sut, "aaa", t)

	sut = NewQueueSlice("aaa", "bbb", 123)
	enqueueTest(sut, t)
}

func TestQueueLinkedList(t *testing.T) {

	sut := &QueueLinkedList{}
	lengthTest(sut, 0, t)

	sut = &QueueLinkedList{}
	isEmptyTest(sut, true, t)

	sut = &QueueLinkedList{}
	dequeueTest(sut, nil, t)

	sut = &QueueLinkedList{}
	enqueueTest(sut, t)

	sut = NewQueueLinkedList("aaa", "bbb", 123)
	lengthTest(sut, 3, t)

	sut = NewQueueLinkedList("aaa", "bbb", 123)
	isEmptyTest(sut, false, t)

	sut = NewQueueLinkedList("aaa", "bbb", 123)
	dequeueTest(sut, "aaa", t)

	sut = NewQueueLinkedList("aaa", "bbb", 123)
	enqueueTest(sut, t)
}

func lengthTest(sut Queue, expected int, t *testing.T) {
	if sut.Length() != expected {
		t.Errorf("Lengthが期待値と相違")
	}
}

func isEmptyTest(sut Queue, expected bool, t *testing.T) {
	if sut.IsEmpty() != expected {
		t.Errorf("IsEmptyが期待値と相違")
	}
}

func enqueueTest(sut Queue, t *testing.T) {

	expectedLength := sut.Length() + 1
	sut.Enqueue(1)

	if sut.Length() != expectedLength {
		t.Errorf("Enqueue後のlengthが期待値と相違")
	}
}

func dequeueTest(sut Queue, expectedDequeued interface{}, t *testing.T) {

	var expectedLength int
	if sut.Length() == 0 {
		expectedLength = 0
	} else {
		expectedLength = sut.Length() - 1
	}

	actual := sut.Dequeue()

	if sut.Length() != expectedLength {
		t.Errorf("Dequeue後のlengthが期待値と相違")
	}
	if actual != expectedDequeued {
		t.Errorf("Dequeueした値が期待値と相違")
	}
}
