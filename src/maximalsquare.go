package main

import "math"

func maximalSquare(matrix [][]string) int {
	rowSize := len(matrix)
	if rowSize == 0 {
		return 0
	}

	colSize := len(matrix[0])
	maxSideDP := make([][]int, rowSize)
	maxSide := 0

	for r := 0; r < rowSize; r++ {
		maxSideDP[r] = make([]int, colSize)
		for c := 0; c < colSize; c++ {
			if matrix[r][c] == "1" {
				maxSideDP[r][c] = 1

				neighbourMinVal := math.MaxInt
				if r < 1 || c < 1 {
					neighbourMinVal = 0
				} else {
					if maxSideDP[r-1][c] < neighbourMinVal {
						neighbourMinVal = maxSideDP[r-1][c]
					}
					if maxSideDP[r][c-1] < neighbourMinVal {
						neighbourMinVal = maxSideDP[r][c-1]
					}
					if maxSideDP[r-1][c-1] < neighbourMinVal {
						neighbourMinVal = maxSideDP[r-1][c-1]
					}
				}

				if neighbourMinVal != math.MaxInt {
					maxSideDP[r][c] += neighbourMinVal
				}

				if maxSideDP[r][c] > maxSide {
					maxSide = maxSideDP[r][c]
				}
			}
		}
	}

	return maxSide * maxSide
}
