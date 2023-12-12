package main

func dutchNationalFlag(arr []int64, index int32) {
	pivot := arr[index]

	small := 0
	equal := 0
	large := len(arr) - 1

	for equal <= large {
		if arr[equal] < pivot {
			temp := arr[small]
			arr[small] = arr[equal]
			arr[equal] = temp
			small += 1
			equal += 1
		} else if arr[equal] == pivot {
			equal += 1
		} else {
			temp := arr[large]
			arr[large] = arr[equal]
			arr[equal] = temp
			large -= 1
		}
	}
}
