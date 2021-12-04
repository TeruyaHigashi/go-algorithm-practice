package stack

import (
	"testing"
)

func TestStackSlice(t *testing.T) {

	sut := &StackSlice{}
	lengthTest(sut, 0, t)

	sut = &StackSlice{}
	isEmptyTest(sut, true, t)

	sut = &StackSlice{}
	peekTest(sut, nil, t)

	sut = &StackSlice{}
	pushTest(sut, t)

	sut = &StackSlice{}
	popTest(sut, nil, t)

	sut = NewStackSlice("aaa", "bbb", 123)
	lengthTest(sut, 3, t)

	sut = NewStackSlice("aaa", "bbb", 123)
	isEmptyTest(sut, false, t)

	sut = NewStackSlice("aaa", "bbb", 123)
	peekTest(sut, 123, t)

	sut = NewStackSlice("aaa", "bbb", 123)
	pushTest(sut, t)

	sut = NewStackSlice("aaa", "bbb", 123)
	popTest(sut, 123, t)
}

func TestStackLinkedList(t *testing.T) {

	sut := &StackLinkedList{}
	lengthTest(sut, 0, t)

	sut = &StackLinkedList{}
	isEmptyTest(sut, true, t)

	sut = &StackLinkedList{}
	peekTest(sut, nil, t)

	sut = &StackLinkedList{}
	pushTest(sut, t)

	sut = &StackLinkedList{}
	popTest(sut, nil, t)

	sut = NewStackLinkedList("aaa", "bbb", 123)
	lengthTest(sut, 3, t)

	sut = NewStackLinkedList("aaa", "bbb", 123)
	isEmptyTest(sut, false, t)

	sut = NewStackLinkedList("aaa", "bbb", 123)
	peekTest(sut, 123, t)

	sut = NewStackLinkedList("aaa", "bbb", 123)
	pushTest(sut, t)

	sut = NewStackLinkedList("aaa", "bbb", 123)
	popTest(sut, 123, t)
}

func lengthTest(sut Stack, expected int, t *testing.T) {
	if sut.Length() != expected {
		t.Errorf("Lengthが期待値と相違")
	}
}

func isEmptyTest(sut Stack, expected bool, t *testing.T) {
	if sut.IsEmpty() != expected {
		t.Errorf("IsEmptyが期待値と相違")
	}
}

func peekTest(sut Stack, expected interface{}, t *testing.T) {
	if sut.Peek() != expected {
		t.Errorf("Peekが期待値と相違")
	}
}

func pushTest(sut Stack, t *testing.T) {

	expectedLength := sut.Length() + 1
	sut.Push(1)

	if sut.Length() != expectedLength {
		t.Errorf("Push後のlengthが期待値と相違")
	}
	if sut.Peek() != 1 {
		t.Errorf("Push後のtopが期待値と相違")
	}
}

func popTest(sut Stack, expectedPoped interface{}, t *testing.T) {

	var expectedLength int
	if sut.Length() == 0 {
		expectedLength = 0
	} else {
		expectedLength = sut.Length() - 1
	}

	actual := sut.Pop()

	if sut.Length() != expectedLength {
		t.Errorf("Pop後のlengthが期待値と相違")
	}
	if actual != expectedPoped {
		t.Errorf("Popした値が期待値と相違")
	}
}
