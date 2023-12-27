package stack

type Stack struct {
	data []int
	size int
}

func NewStack() *Stack {
	return &Stack{
		data: []int{},
		size: 0,
	}
}

func (s *Stack) push(e int) {
	s.data = append(s.data, e)
	s.size += 1
}

func (s *Stack) pop() int {
	if s.size == 0 {
		return -1
	}

	e := s.data[s.size-1]
	s.data = s.data[0 : s.size-1]
	s.size -= 1
	return e
}

func (s *Stack) peek() int {
	if s.size == 0 {
		return -1
	}

	return s.data[s.size-1]
}

func (s *Stack) isEmpty() bool {
	return s.size <= 0
}
