package main

func firstOccurenceOfK(sortedArr []int, k int) int {
	l := 0
	r := len(sortedArr) - 1
	firstOccerenceIndex := -1
	for l <= r {
		m := l + (r-l)/2
		if sortedArr[m] < k {
			l = m + 1
		} else if sortedArr[m] > k {
			r = m - 1
		} else {
			firstOccerenceIndex = m
			r = m - 1
		}
	}

	return firstOccerenceIndex
}
