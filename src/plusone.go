package main

func plusOne(arr []int32) []int32 {
	arrLen := len(arr)
	arr[arrLen-1] += 1
	for i := arrLen - 1; i > 0 && arr[i] == 10; i-- {
		arr[i] = 0
		arr[i-1] += 1
	}

	if arr[0] == 10 {
		arr[0] = 0
		return append([]int32{1}, arr...)
	}

	return arr
}
