package stack

import "testing"

func TestStackSlice(t *testing.T) {

	stack := &StackSlice{}
	lengthTest(stack, 0, t)
	isEmptyTest(stack, true, t)
	pushTest(stack, t)
	popTest(stack, t)

	stack = stack.NewStackSlice("aaa", "bbb", 123)
	lengthTest(stack, 3, t)
	peekTest(stack, 123, t)
	isEmptyTest(stack, false, t)
	pushTest(stack, t)
	popTest(stack, t)
}

func TestStackLinkedList(t *testing.T) {

	stack := &StackLinkedList{}
	lengthTest(stack, 0, t)
	isEmptyTest(stack, true, t)
	pushTest(stack, t)
	popTest(stack, t)

	stack = stack.NewStackLinkedList("aaa", "bbb", 123)
	lengthTest(stack, 3, t)
	peekTest(stack, 123, t)
	isEmptyTest(stack, false, t)
	pushTest(stack, t)
	popTest(stack, t)
}

func lengthTest(stack Stack, expected int, t *testing.T) {
	if stack.Length() != expected {
		t.Errorf("Lengthが異常")
	}
}

func isEmptyTest(stack Stack, expected bool, t *testing.T) {
	if stack.IsEmpty() != expected {
		t.Errorf("IsEmptyが異常")
	}
}

func peekTest(stack Stack, expected interface{}, t *testing.T) {
	if stack.Peek() != expected {
		t.Errorf("Peekが異常")
	}
}

func pushTest(stack Stack, t *testing.T) {

	expectedLength := stack.Length() + 1
	stack.Push(1)

	if stack.Length() != expectedLength {
		t.Errorf("Push後のlengthが異常")
	}
	if stack.Peek() != 1 {
		t.Errorf("Push後のtopが異常")
	}
}

func popTest(stack Stack, t *testing.T) {

	expectedLength := stack.Length() - 1
	expectedPoped := stack.Peek()
	actual := stack.Pop()

	if stack.Length() != expectedLength {
		t.Errorf("Pop後のlengthが異常")
	}
	if actual != expectedPoped {
		t.Errorf("Popした値が異常")
	}
}
