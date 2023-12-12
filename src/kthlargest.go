package main

func findKthLargestWithInPlaceOperations(arr []int, k int) int {
	arrLen := len(arr)

	if k > arrLen {
		return -1
	}

	left := 0
	right := arrLen - 1

	for left <= right {
		pivot := left + (right-left)/2
		newPivotIndex := partitionAroundPivot(arr, left, right, pivot)
		if newPivotIndex == arrLen-k {
			return arr[newPivotIndex]
		} else if newPivotIndex < arrLen-k {
			left = newPivotIndex + 1
		} else {
			right = newPivotIndex - 1
		}
	}

	return -1
}

func partitionAroundPivot(arr []int, left int, right int, pivot int) int {
	newPivotIndex := left
	pivotVal := arr[pivot]
	for left < right {
		if arr[left] <= pivotVal {
			temp := arr[newPivotIndex]
			arr[newPivotIndex] = arr[left]
			arr[left] = temp
			newPivotIndex += 1
		}
		left += 1
	}

	temp := arr[newPivotIndex]
	arr[newPivotIndex] = pivotVal
	arr[pivot] = temp

	return newPivotIndex
}
