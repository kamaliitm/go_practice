package main

func searchInsert(nums []int, target int) int {
	arrLen := len(nums)
	targetIndex := arrLen
	i := 0
	j := arrLen - 1

	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			targetIndex = mid
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	return targetIndex
}
