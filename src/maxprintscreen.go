package main

func printMaxA(n int32) int32 {
	dp := make([]int32, n+1)

	for i := int32(0); i <= n; i++ {
		if i <= 6 {
			dp[i] = i
		} else {
			maxLen := int32(0)
			if dp[i-3]*2 > int32(maxLen) {
				maxLen = dp[i-3] * 2
			}
			if dp[i-4]*3 > maxLen {
				maxLen = dp[i-4] * 3
			}
			if dp[i-5] > maxLen {
				maxLen = dp[i-5] * 4
			}

			dp[i] = maxLen

		}

	}

	return dp[n]
}
