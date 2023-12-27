package main

func rotate(nums []int, k int) {
	arrLen := len(nums)
	if arrLen == 0 || k == 0 {
		return
	}

	reverseNums(nums, 0, arrLen-1)
	reverseNums(nums, 0, k-1)
	reverseNums(nums, k, arrLen-1)
}

func reverseNums(nums []int, start int, end int) {
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}
}
