package stack

type StackLinkedList struct {
	top    *node
	length int
}

type node struct {
	value *interface{}
	next  *node
}

func (s *StackLinkedList) NewStackLinkedList(elems ...interface{}) *StackLinkedList {
	if elems == nil {
		return &StackLinkedList{}
	}
	stack := StackLinkedList{}
	for _, elem := range elems {
		stack.Push(elem)
	}
	return &stack
}

func (s *StackLinkedList) Push(elem interface{}) {
	pushed := node{value: &elem, next: nil}
	if s.length > 0 {
		pushed.next = s.top
	}
	s.top = &pushed
	s.length++
}

func (s *StackLinkedList) Pop() interface{} {
	poped := *(s.top.value)
	if s.Length() == 1 {
		s.top = nil
	} else {
		s.top = s.top.next
	}
	s.length--
	return poped
}

func (s *StackLinkedList) Peek() interface{} {
	return *(s.top.value)
}

func (s *StackLinkedList) Length() int {
	return s.length
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.Length() == 0
}
