package main

type RuneStack struct {
	data []rune
	size int64
}

func (stack *RuneStack) push(r rune) {
	if stack.data == nil {
		stack.data = []rune{}
	}

	stack.data = append(stack.data, r)
	stack.size += 1
}

func (stack *RuneStack) pop() rune {
	if stack.data == nil || stack.size == 0 {
		return ' '
	}

	r := stack.data[stack.size-1]
	stack.data = stack.data[:stack.size-1]
	stack.size -= 1

	return r
}

func (stack *RuneStack) peek() rune {
	if stack.data == nil || stack.size == 0 {
		return ' '
	}

	return stack.data[stack.size-1]
}

func isBalanced(s string) string {
	// Write your code here
	if len(s) == 0 {
		return "YES"
	}

	stack := RuneStack{}

	for _, r := range s {
		switch r {
		case '{', '(', '[':
			stack.push(r)
		case '}':
			if stack.peek() != '{' {
				return "NO"
			}
			stack.pop()
		case ')':
			if stack.peek() != '(' {
				return "NO"
			}
			stack.pop()
		case ']':
			if stack.peek() != '[' {
				return "NO"
			}
			stack.pop()
		default:
			return "NO"
		}
	}

	if stack.size == 0 {
		return "YES"
	}

	return "NO"
}
