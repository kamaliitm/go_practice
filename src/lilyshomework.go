package main

import (
	"sort"
)

func lilysHomework(arr []int32) int32 {

	numSwapsAscOrder := numSwaps(arr, true)
	numSwapsDescOrder := numSwaps(arr, false)

	if numSwapsAscOrder < numSwapsDescOrder {
		return numSwapsAscOrder
	}

	return numSwapsDescOrder
}

func numSwaps(arr []int32, sortOrderAsc bool) int32 {
	tempArr := []int32{}
	tempArr = append(tempArr, arr...)
	sortedArr := []int32{}
	sortedArr = append(sortedArr, arr...)
	sort.Slice(sortedArr, func(i, j int) bool {
		if sortOrderAsc {
			return sortedArr[i] < sortedArr[j]
		}
		return sortedArr[i] > sortedArr[j] // sort order descending
	})

	valueToIndexMap := map[int32]int{}
	for index, val := range tempArr {
		valueToIndexMap[val] = index
	}

	numSwaps := 0
	for i, sortedVal := range sortedArr {
		if sortedVal != tempArr[i] {
			indexInOriginalArr := valueToIndexMap[sortedVal]
			tempArr[i], tempArr[indexInOriginalArr] = sortedVal, tempArr[i]
			valueToIndexMap[sortedVal] = i
			valueToIndexMap[tempArr[indexInOriginalArr]] = indexInOriginalArr
			numSwaps += 1
		}
	}

	// fmt.Println(tempArr)
	// fmt.Println(numSwaps)

	return int32(numSwaps)
}
