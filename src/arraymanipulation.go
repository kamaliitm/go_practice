package main

func at56rrayManipulation(n int32, queries [][]int32) int64 {

	data := make([]int32, n)

	for _, query := range queries {
		num := query[2]
		data[query[0]-1] += num
		if query[1] < n {
			data[query[1]] -= num
		}
	}

	max := int64(0)
	totalSum := int64(0)

	for _, elem := range data {
		totalSum += int64(elem)
		if totalSum > max {
			max = totalSum
		}
	}

	return max
}
