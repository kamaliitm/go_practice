package main

type ParenthesisStack struct {
	data []int
	size int
}

func NewParenthesisStack() *ParenthesisStack {
	return &ParenthesisStack{
		data: []int{},
		size: 0,
	}
}

func (s *ParenthesisStack) push(e int) {
	s.data = append(s.data, e)
	s.size += 1
}

func (s *ParenthesisStack) pop() int {
	if s.size == 0 {
		return -1
	}

	e := s.data[s.size-1]
	s.data = s.data[0 : s.size-1]
	s.size -= 1
	return e
}

func (s *ParenthesisStack) peek() int {
	if s.size == 0 {
		return -1
	}

	return s.data[s.size-1]
}

func (s *ParenthesisStack) isEmpty() bool {
	return s.size <= 0
}

func longestValidParentheses(s string) int {
	maxLen := 0
	currLen := 0

	stack := &ParenthesisStack{}

	for i, r := range s {
		if r == '(' {
			stack.push(i)
		} else {
			if !stack.isEmpty() {
				currLen += 2
				stack.pop()
			} else {
				if currLen > maxLen {
					maxLen = currLen
				}
				currLen = 0
			}
		}
	}

	if currLen > maxLen {
		maxLen = currLen
	}

	return maxLen
}
