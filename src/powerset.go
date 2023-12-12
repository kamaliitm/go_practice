package main

func powerSet(arr []int32) [][]int32 {
	res := [][]int32{}

	powerSetRecursive(arr, &res)

	return res
}

func powerSetRecursive(arr []int32, res *[][]int32) {
	if len(arr) == 0 {
		return
	}

	tempArr := []int32{}
	for _, entry := range arr {
		tempArr = append(tempArr, entry)
		*res = append(*res, tempArr)
	}

	if len(arr) > 1 {
		powerSetRecursive(arr[1:], res)
	}
}
