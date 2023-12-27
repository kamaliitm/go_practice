package main

func partitionLabels(s string) []int {
	runes := []rune(s)
	maxReachMap := map[rune]int{}

	for i, r := range runes {
		maxReachMap[r] = i
	}

	partitions := []int{}
	i := 0
	strLen := len(runes)
	for i < strLen {
		maxReach := maxReachMap[runes[i]]
		for j := i + 1; j < maxReach; j++ {
			maxReachForJ := maxReachMap[runes[j]]
			if maxReachForJ > maxReach {
				maxReach = maxReachForJ
			}
		}

		partitions = append(partitions, maxReach-i+1)
		i = maxReach + 1
	}

	return partitions
}
