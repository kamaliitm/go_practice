package main

func canPartition(nums []int) bool {
	numLen := len(nums)
	if numLen == 0 {
		return false
	}

	sum := 0
	for i := 0; i < numLen; i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	sum = sum / 2
	dp := make([][]bool, numLen+1)
	dp[0] = make([]bool, sum+1)
	for i := 1; i <= numLen; i++ {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
		for j := 1; j <= sum; j++ {
			dp[0][j] = false
			withoutThis := dp[i-1][j]
			withThis := false
			if j >= nums[i-1] {
				withThis = dp[i-1][j-nums[i-1]]
			}

			dp[i][j] = withThis || withoutThis
		}
	}

	return dp[numLen][sum]

}
