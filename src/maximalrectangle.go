package main

type MaxRectStackNode struct {
	height int
	index  int
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rowSize := len(matrix)
	colSize := len(matrix[0])
	maxArea := 0

	histArrForRow := make([]int, colSize)
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if string(matrix[i][j]) == "1" {
				histArrForRow[j] += 1
			} else {
				histArrForRow[j] = 0
			}
		}

		maxAreaFromRow := maxHistogramAreaForRow(histArrForRow)
		if maxArea < maxAreaFromRow {
			maxArea = maxAreaFromRow
		}
	}

	return maxArea
}

func maxHistogramAreaForRow(arr []int) int {
	maxArea := 0
	stack := []MaxRectStackNode{}
	stackSize := 0

	for i, h := range arr {
		extendedIndex := i
		for stackSize > 0 && stack[stackSize-1].height > h {
			node := stack[stackSize-1]
			areaFromNode := node.height * (i - node.index)
			if maxArea < areaFromNode {
				maxArea = areaFromNode
			}
			extendedIndex = node.index
			stackSize -= 1
			stack = stack[:stackSize]
		}

		stack = append(stack, MaxRectStackNode{height: h, index: extendedIndex})
		stackSize += 1
	}

	arrLen := len(arr)
	for stackSize > 0 {
		node := stack[stackSize-1]
		areaFromNode := node.height * (arrLen - node.index)
		if maxArea < areaFromNode {
			maxArea = areaFromNode
		}
		stackSize -= 1
	}

	return maxArea
}
