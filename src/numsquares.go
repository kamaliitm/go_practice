package main

import "math"

func numSquares(n int) int {

	squaresDp := make([]int, n+1)
	squaresDp[0] = 1

	for i := 1; i <= n; i++ {
		start := 1
		end := int(math.Sqrt(float64(i)))
		minNumSquares := math.MaxInt
		for j := start; j <= end; j++ {
			sq := j * j
			if i >= sq {
				res := 1 + squaresDp[i-sq]
				if res < minNumSquares {
					minNumSquares = res
				}
			}
		}
		squaresDp[i] = minNumSquares
	}

	return squaresDp[n] - 1
}
