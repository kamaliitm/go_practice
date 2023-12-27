package main

type TemperatureStackNode struct {
	temperature int
	index       int
}

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))

	stack := []TemperatureStackNode{}
	stackSize := 0
	for i, t := range temperatures {
		for stackSize > 0 && stack[stackSize-1].temperature < t {
			indexToUpdate := stack[stackSize-1].index
			answer[indexToUpdate] = i - indexToUpdate
			stack = stack[:stackSize-1]
			stackSize -= 1
		}

		stack = append(stack, TemperatureStackNode{temperature: t, index: i})
		stackSize += 1
	}

	return answer
}
