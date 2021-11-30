package stack

type StackSlice struct {
	list   []interface{}
	length int
}

func (s *StackSlice) NewStackSlice(elems ...interface{}) *StackSlice {
	for _, elem := range elems {
		s.Push(elem)
	}
	return s
}

func (s *StackSlice) Push(elem interface{}) {
	s.list = append(s.list, elem)
	s.length++
}

func (s *StackSlice) Pop() interface{} {
	poped := s.list[s.length-1]
	s.list = s.list[:s.length-1]
	s.length--
	return poped
}

func (s *StackSlice) Peek() interface{} {
	return s.list[s.length-1]
}

func (s *StackSlice) Length() int {
	return s.length
}

func (s *StackSlice) IsEmpty() bool {
	return s.Length() == 0
}
