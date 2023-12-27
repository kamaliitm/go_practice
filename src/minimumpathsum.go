package main

import "math"

func minPathSum(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rowSize := len(grid)
	colSize := len(grid[0])

	pathSumDp := make([][]int, rowSize)
	pathSumDp[0] = make([]int, colSize)
	pathSumDp[0][0] = grid[0][0]

	for i := 0; i < rowSize; i++ {
		if i > 0 {
			pathSumDp[i] = make([]int, colSize)
		}

		for j := 0; j < colSize; j++ {
			pathSumDp[i][j] = grid[i][j]

			minSum := math.MaxInt
			if j >= 1 {
				minSum = pathSumDp[i][j-1]
			}

			if i >= 1 && minSum > pathSumDp[i-1][j] {
				minSum = pathSumDp[i-1][j]
			}

			if minSum != math.MaxInt {
				pathSumDp[i][j] += minSum
			}
		}
	}

	return pathSumDp[rowSize-1][colSize-1]

}
