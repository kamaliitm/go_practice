package main

func intersectionOfSortedArrays(arr1 []int32, arr2 []int32) []int32 {
	res := []int32{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			res = append(res, arr1[i])
			i = incrementIndex(arr1, i)
			j = incrementIndex(arr2, j)
		} else if arr1[i] < arr2[j] {
			i = incrementIndex(arr1, i)
		} else {
			j = incrementIndex(arr2, j)
		}
	}

	return res
}

func incrementIndex(arr []int32, currIndex int) int {
	currIndex += 1
	for currIndex < len(arr)-1 {
		if arr[currIndex] == arr[currIndex-1] {
			currIndex += 1
		} else {
			break
		}
	}

	return currIndex
}
