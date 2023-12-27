package main

func numTrees(n int) int {
	numOfTrees := make([]int, n+1)
	numOfTrees[0] = 1
	numOfTrees[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			numOfTrees[i] += numOfTrees[i-j] * numOfTrees[j-1]
		}
	}

	return numOfTrees[n]
}
