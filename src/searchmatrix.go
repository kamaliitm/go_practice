package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rowSize := len(matrix)
	colSize := len(matrix[0])
	for i := 0; i < rowSize; i++ {
		if target < matrix[i][0] {
			return false
		}

		if target > matrix[i][colSize-1] {
			continue
		}

		if binarySearch(matrix[i], target) {
			return true
		}
	}

	return false
}

func binarySearch(arr []int, target int) bool {
	i := 0
	j := len(arr) - 1

	for i <= j {
		mid := i + (j-i)/2
		if target == arr[mid] {
			return true
		} else if target < arr[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	return false
}
