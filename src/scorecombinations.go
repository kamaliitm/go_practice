package main

func findNumOfScoreCombinations(individualScores []int32, finalScore int32) int32 {

	individualScoresLen := len(individualScores)
	numberOfCombinations := make([][]int32, individualScoresLen)

	for i := 0; i < individualScoresLen; i++ {

		numberOfCombinations[i] = make([]int32, finalScore+1)
		numberOfCombinations[i][0] = 1
		currIndividualScore := individualScores[i]

		for s := int32(1); s <= finalScore; s++ {
			withOutCurrScore := int32(0)
			if i > 0 {
				withOutCurrScore = numberOfCombinations[i-1][s]
			}

			withCurrScore := int32(0)
			if s >= currIndividualScore {
				withCurrScore = numberOfCombinations[i][s-currIndividualScore]
			}

			numberOfCombinations[i][s] = withOutCurrScore + withCurrScore
		}
	}

	return numberOfCombinations[individualScoresLen-1][finalScore]
}
