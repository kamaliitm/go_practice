package main

// type Stack struct {
// 	elements        []Element
// 	maxValues       []Element
// 	length          int32
// 	maxValuesLength int32
// }

// type Element struct {
// 	val   int32
// 	count int32
// }

// func (stack *Stack) push(v int32) {
// 	stack.elements = append(stack.elements, Element{val: v})
// 	stack.length += 1

// 	if stack.maxValuesLength == 0 {
// 		stack.maxValues = append(stack.maxValues, Element{val: v, count: 1})
// 		stack.maxValuesLength = 1
// 	} else {
// 		currentMaxElem := &stack.maxValues[stack.maxValuesLength-1]
// 		if currentMaxElem.val < v {
// 			stack.maxValues = append(stack.maxValues, Element{val: v, count: 1})
// 			stack.maxValuesLength += 1
// 		} else if currentMaxElem.val == v {
// 			currentMaxElem.count += 1
// 		}
// 	}
// }

// func (stack *Stack) pop() int32 {
// 	if len(stack.elements) == 0 {
// 		return -1
// 	}

// 	elemToPop := stack.elements[stack.length-1]
// 	currentMaxElem := &stack.maxValues[stack.maxValuesLength-1]
// 	if elemToPop.val == currentMaxElem.val {
// 		if currentMaxElem.count == 1 {
// 			stack.maxValues = stack.maxValues[:stack.maxValuesLength-1]
// 			stack.maxValuesLength -= 1
// 		} else {
// 			currentMaxElem.count -= 1
// 		}
// 	}

// 	stack.elements = stack.elements[:stack.length-1]
// 	stack.length -= 1

// 	return elemToPop.val
// }

// func (stack *Stack) max() int32 {
// 	if len(stack.elements) == 0 {
// 		return -1
// 	}

// 	return stack.maxValues[stack.maxValuesLength-1].val
// }
