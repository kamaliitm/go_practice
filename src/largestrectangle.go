package main

func largestRectangle(h []int32) int64 {
	i := 0
	j := len(h) - 1
	maxArea := int64(0)

	for i < j {
		height := int64(h[i])
		if int64(h[j]) < height {
			height = int64(h[j])
		}

		area := height * int64(j-i+1)
		if area > maxArea {
			maxArea = area
		}

		if h[i] < h[j] {
			i += 1
		} else if h[i] > h[j] {
			j -= 1
		} else {
			i += 1
			j -= 1
		}
	}

	return maxArea
}
