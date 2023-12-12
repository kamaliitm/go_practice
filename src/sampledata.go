package main

import (
	"math/rand"
	"time"
)

func sample(arr []int32, size int32) []int32 {
	rand.Seed(time.Now().UnixNano())
	// min := 10
	// max := 20
	// fmt.Println(rand.Intn(max-min+1) + min)

	arrLen := len(arr)
	for i := 0; i < int(size); i++ {
		randIndex := rand.Intn(arrLen-i) + i
		temp := arr[i]
		arr[i] = arr[randIndex]
		arr[randIndex] = temp
	}

	return arr
}
