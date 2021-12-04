package stack

type StackLinkedList struct {
	top    *node
	length int
}

type node struct {
	value *interface{}
	next  *node
}

func NewStackLinkedList(elems ...interface{}) *StackLinkedList {
	stack := &StackLinkedList{}
	if elems == nil {
		return stack
	}
	for _, elem := range elems {
		stack.Push(elem)
	}
	return stack
}

func (s *StackLinkedList) Push(elem interface{}) {
	pushed := &node{value: &elem, next: nil}
	if s.length > 0 {
		pushed.next = s.top
	}
	s.top = pushed
	s.length++
}

func (s *StackLinkedList) Pop() interface{} {
	if s.Length() == 0 {
		return nil
	}
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
	if s.Length() == 0 {
		return nil
	}
	return *(s.top.value)
}

func (s *StackLinkedList) Length() int {
	return s.length
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.Length() == 0
}
