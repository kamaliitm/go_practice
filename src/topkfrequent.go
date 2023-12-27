package main

func topKFrequent(nums []int, k int) []int {
	frequencyMap := map[int]int{}
	for _, num := range nums {
		if freq, ok := frequencyMap[num]; ok {
			frequencyMap[num] = freq + 1
		} else {
			frequencyMap[num] = 1
		}
	}

	numsLen := len(nums)
	buckets := make([][]int, numsLen+1)
	for num, freq := range frequencyMap {
		if len(buckets[freq]) == 0 {
			buckets[freq] = []int{num}
		} else {
			buckets[freq] = append(buckets[freq], num)
		}
	}

	topK := make([]int, k)
	i := 0
	for j := numsLen; j >= 0 && i < k; j-- {
		if len(buckets[j]) > 0 {
			for _, num := range buckets[j] {
				topK[i] = num
				i += 1
				if i == k {
					break
				}
			}
		}
		// i += 1
	}

	return topK
}
