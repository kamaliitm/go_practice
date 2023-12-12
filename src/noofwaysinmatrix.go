package main

func findNumOfWaysToBottomRightInMatrix(noOfRows int32, noOfCols int32) int32 {

	numOfWays := make([][]int32, noOfRows)
	for i := 0; i < int(noOfRows); i++ {
		numOfWays[i] = make([]int32, noOfCols)
	}

	numOfWays[0][0] = 1
	for i := int32(0); i < noOfRows; i++ {
		for j := int32(0); j < noOfCols; j++ {
			if i == 0 && j == 0 {
				continue
			}
			fromLeft := int32(0)
			if j > 0 {
				fromLeft = numOfWays[i][j-1]
			}

			fromTop := int32(0)
			if i > 0 {
				fromTop = numOfWays[i-1][j]
			}

			numOfWays[i][j] = fromLeft + fromTop
		}
	}

	return numOfWays[noOfRows-1][noOfCols-1]
}
