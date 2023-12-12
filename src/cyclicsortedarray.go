package main

// assumption: all elements must be distinct
func findSmallestInCyclicSortedArray(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	l := 0
	r := len(arr) - 1
	smallest := arr[0]

	for l <= r {
		m := l + (r-l)/2
		if arr[m] < smallest {
			smallest = arr[m]
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return smallest

}
