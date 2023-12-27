package main

import "math"

func jump(nums []int) int {
	numLen := len(nums)
	steps := make([]int, numLen)
	steps[0] = 0

	for i := 1; i < numLen; i++ {
		steps[i] = -1
	}

	getSteps(nums, numLen-1, steps)

	return steps[numLen-1]
}

func getSteps(nums []int, index int, steps []int) int {
	if index < 0 {
		return math.MaxInt
	}

	if steps[index] == -1 {
		minSteps := math.MaxInt
		for i := index - 1; i >= 0; i-- {
			if nums[i] >= index-i {
				stepsToI := getSteps(nums, i, steps)
				if minSteps > stepsToI+1 {
					minSteps = stepsToI + 1
				}
			}
		}
		steps[index] = minSteps
	}

	return steps[index]

}
