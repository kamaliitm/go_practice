package main

func deleteDuplicates(arr []int32) []int32 {
	if len(arr) == 0 {
		return arr
	}

	writeIndex := 1
	for i := 1; i < len(arr); i++ {
		if arr[writeIndex-1] != arr[i] {
			arr[writeIndex] = arr[i]
			writeIndex += 1
		}
	}

	return arr
}
