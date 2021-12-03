package stack

import (
	"testing"
)

func TestStackSlice(t *testing.T) {

	stack := &StackSlice{}
	lengthTest(stack, 0, t)

	stack = &StackSlice{}
	isEmptyTest(stack, true, t)

	stack = &StackSlice{}
	pushTest(stack, t)

	stack = NewStackSlice("aaa", "bbb", 123)
	lengthTest(stack, 3, t)

	stack = NewStackSlice("aaa", "bbb", 123)
	isEmptyTest(stack, false, t)

	stack = NewStackSlice("aaa", "bbb", 123)
	peekTest(stack, 123, t)

	stack = NewStackSlice("aaa", "bbb", 123)
	popTest(stack, t)

	stack = NewStackSlice("aaa", "bbb", 123)
	pushTest(stack, t)
}

func TestStackLinkedList(t *testing.T) {

	stack := &StackLinkedList{}
	lengthTest(stack, 0, t)

	stack = &StackLinkedList{}
	isEmptyTest(stack, true, t)

	stack = &StackLinkedList{}
	pushTest(stack, t)

	stack = NewStackLinkedList("aaa", "bbb", 123)
	lengthTest(stack, 3, t)

	stack = NewStackLinkedList("aaa", "bbb", 123)
	isEmptyTest(stack, false, t)

	stack = NewStackLinkedList("aaa", "bbb", 123)
	peekTest(stack, 123, t)

	stack = NewStackLinkedList("aaa", "bbb", 123)
	popTest(stack, t)

	stack = NewStackLinkedList("aaa", "bbb", 123)
	pushTest(stack, t)
}

func lengthTest(stack Stack, expected int, t *testing.T) {
	if stack.Length() != expected {
		t.Errorf("Lengthが期待値と相違")
	}
}

func isEmptyTest(stack Stack, expected bool, t *testing.T) {
	if stack.IsEmpty() != expected {
		t.Errorf("IsEmptyが期待値と相違")
	}
}

func peekTest(stack Stack, expected interface{}, t *testing.T) {
	if stack.Peek() != expected {
		t.Errorf("Peekが期待値と相違")
	}
}

func pushTest(stack Stack, t *testing.T) {

	expectedLength := stack.Length() + 1
	stack.Push(1)

	if stack.Length() != expectedLength {
		t.Errorf("Push後のlengthが期待値と相違")
	}
	if stack.Peek() != 1 {
		t.Errorf("Push後のtopが期待値と相違")
	}
}

func popTest(stack Stack, t *testing.T) {

	expectedLength := stack.Length() - 1
	expectedPoped := stack.Peek()
	actual := stack.Pop()

	if stack.Length() != expectedLength {
		t.Errorf("Pop後のlengthが期待値と相違")
	}
	if actual != expectedPoped {
		t.Errorf("Popした値が期待値と相違")
	}
}
