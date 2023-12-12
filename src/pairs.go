package main

import "sort"

func pairs(k int32, arr []int32) int32 {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	arrLen := len(arr)
	if arrLen == 0 {
		return 0
	}

	numOfPairs := int32(0)
	i := 0
	j := 1
	for i < arrLen && j < arrLen {
		diff := arr[j] - arr[i]
		if diff == k {
			numOfPairs += 1
			j += 1
		} else if diff < k {
			j += 1
		} else {
			i += 1
		}
	}

	return numOfPairs
}
