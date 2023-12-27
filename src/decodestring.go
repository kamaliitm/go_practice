package main

import "strconv"

type DecodeStrElem struct {
	char  rune
	count int
}

type DecodeStrStack struct {
	data []*DecodeStrElem
	size int
}

func NewStack() *DecodeStrStack {
	return &DecodeStrStack{
		data: []*DecodeStrElem{},
		size: 0,
	}
}

func (s *DecodeStrStack) push(e *DecodeStrElem) {
	s.data = append(s.data, e)
	s.size += 1
}

func (s *DecodeStrStack) pop() *DecodeStrElem {
	if s.size == 0 {
		return nil
	}

	e := s.data[s.size-1]
	s.data = s.data[0 : s.size-1]
	s.size -= 1
	return e
}

func (s *DecodeStrStack) peek() *DecodeStrElem {
	if s.size == 0 {
		return nil
	}

	return s.data[s.size-1]
}

func (s *DecodeStrStack) isEmpty() bool {
	return s.size <= 0
}

func decodeString(s string) string {

	stack := &DecodeStrStack{}
	currCount := 0
	countStr := []rune{}
	decodedStr := ""
	// tempStr := ""

	for _, r := range s {

		if r >= '0' && r <= '9' {
			countStr = append(countStr, r)
			continue
		} else {
			if r == '[' {
				var err error
				currCount, err = strconv.Atoi(string(countStr))
				if err != nil {
					panic("invalid integer")
				}

				countStr = []rune{}
			} else if r == ']' {
				tempStr := ""
				count := 0
				for stack.peek() != nil {
					elem := stack.pop()
					tempStr = string(elem.char) + tempStr
					if elem.count > 0 {
						count = elem.count
						break
					}
				}

				str := tempStr
				for i := 1; i < count; i++ {
					str += tempStr
				}

				if !stack.isEmpty() {
					for _, ch := range str {
						stack.push(&DecodeStrElem{char: ch})
					}
				} else {
					decodedStr += str
				}

			} else {
				stack.push(&DecodeStrElem{char: r, count: currCount})
				currCount = 0
			}

		}
	}

	remainingStr := ""
	for !stack.isEmpty() {
		remainingStr = string(stack.pop().char) + remainingStr
	}

	decodedStr += remainingStr

	return decodedStr
}
