package main

func knapsackIterative(valuesAndWeights [][]int32, maxWeight int32) int32 {
	if len(valuesAndWeights) == 0 {
		return 0
	}

	numOfItems := len(valuesAndWeights)
	maxValuesDP := make([][]int32, numOfItems+1)
	maxValuesDP[0] = make([]int32, maxWeight+1)

	for i := 1; i <= numOfItems; i++ {
		maxValuesDP[i] = make([]int32, maxWeight+1)
		maxValuesDP[i][0] = 0
		for j := 1; j <= int(maxWeight); j++ {
			maxValuesDP[0][j] = 0

			withoutThisItem := maxValuesDP[i-1][j]

			withThisItem := int32(0)
			if j >= int(valuesAndWeights[i-1][1]) {
				withThisItem = valuesAndWeights[i-1][0] + maxValuesDP[i-1][j-int(valuesAndWeights[i-1][1])]
			}

			maxValuesDP[i][j] = withoutThisItem
			if withThisItem > withoutThisItem {
				maxValuesDP[i][j] = withThisItem
			}
		}
	}

	return maxValuesDP[numOfItems][maxWeight]

}

func knapsack(valuesAndWeights [][]int32, maxWeight int32) int32 {
	numOfItems := len(valuesAndWeights)
	maxValues := make([][]int32, numOfItems)
	for i := range maxValues {
		maxValues[i] = make([]int32, maxWeight+1)
		maxValues[i][0] = 0
		for j := 1; j <= int(maxWeight); j++ {
			maxValues[i][j] = -1
		}
	}

	return findMaxValWithWeightConstraint(valuesAndWeights, numOfItems-1, maxWeight, &maxValues)
}

func findMaxValWithWeightConstraint(valuesAndWeights [][]int32, itemIndex int, maxWeight int32, maxValuesPtr *[][]int32) int32 {
	maxValues := *maxValuesPtr
	if itemIndex < 0 || itemIndex >= len(maxValues) || maxWeight < 0 || maxWeight >= int32(len(maxValues[0])) {
		return 0
	}

	if maxValues[itemIndex][maxWeight] == -1 {
		if valuesAndWeights[itemIndex][1] > maxWeight {
			maxValues[itemIndex][maxWeight] = findMaxValWithWeightConstraint(
				valuesAndWeights, itemIndex-1, maxWeight, maxValuesPtr,
			)
		} else {
			withOutCurrentItem := findMaxValWithWeightConstraint(
				valuesAndWeights, itemIndex-1, maxWeight, maxValuesPtr,
			)
			withCurrentItem := valuesAndWeights[itemIndex][0] + findMaxValWithWeightConstraint(
				valuesAndWeights, itemIndex-1, maxWeight-valuesAndWeights[itemIndex][1], maxValuesPtr,
			)

			maxValue := withOutCurrentItem
			if withCurrentItem > maxValue {
				maxValue = withCurrentItem
			}

			maxValues[itemIndex][maxWeight] = maxValue
		}
	}

	return maxValues[itemIndex][maxWeight]
}
