package main

func waterTrapped(height []int) int {
	maxHeightsFromLeft := make([]int, len(height))
	maxLeft := 0
	for i, h := range height {
		if h > maxLeft {
			maxLeft = h
		}
		maxHeightsFromLeft[i] = maxLeft
	}

	// update max heights based on min of max from left and right
	maxRight := 0
	waterTrapped := 0
	for j := len(height) - 1; j >= 0; j-- {
		if height[j] > maxRight {
			maxRight = height[j]
		}

		if maxRight < maxHeightsFromLeft[j] {
			maxHeightsFromLeft[j] = maxRight
		}

		waterTrapped += maxHeightsFromLeft[j] - height[j]
	}

	return waterTrapped
}
