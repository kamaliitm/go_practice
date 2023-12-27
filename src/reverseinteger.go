package main

import "math"

func reverseInt(x int) int {
	reversed := int32(0)

	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}

	y := int32(x)

	for y > 0 {
		reversed = (reversed * 10) + (y % 10)
		if reversed > math.MaxInt32 || reversed < math.MinInt32 {
			return 0
		}
		y = y / 10
	}

	if isNegative {
		reversed = -reversed
	}

	return int(reversed)
}

func multiplyAndCheckOverflow(a int32, b int32) (int32, bool) {
	mul := a * b

	res := mul / b

	if a != res {
		return mul, true
	}

	return mul, false

}
