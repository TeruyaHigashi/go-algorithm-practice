package stack

type StackSlice struct {
	list   []interface{}
	length int
}

func NewStackSlice(elems ...interface{}) *StackSlice {
	stack := &StackSlice{}
	for _, elem := range elems {
		stack.Push(elem)
	}
	return stack
}

func (s *StackSlice) Push(elem interface{}) {
	s.list = append(s.list, elem)
	s.length++
}

func (s *StackSlice) Pop() interface{} {
	if s.Length() == 0 {
		return nil
	}
	poped := s.list[s.length-1]
	s.list = s.list[:s.length-1]
	s.length--
	return poped
}

func (s *StackSlice) Peek() interface{} {
	if s.Length() == 0 {
		return nil
	}
	return s.list[s.length-1]
}

func (s *StackSlice) Length() int {
	return s.length
}

func (s *StackSlice) IsEmpty() bool {
	return s.Length() == 0
}
