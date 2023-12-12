package main

// func kSortArr(arr []int, k int) []int {
// 	if k <= 0 {
// 		return arr
// 	}

// 	minHeap := &MinHeap{}

// 	arrLen := len(arr)
// 	for i := 0; i <= k && i < arrLen; i++ {
// 		minHeap.insert(HeapElement{arr[i], 0})
// 	}

// 	for i := 0; i < len(arr); i++ {
// 		arr[i] = minHeap.extractMin().val
// 		if i < len(arr)-k-1 {
// 			minHeap.insert(HeapElement{arr[i+k+1], 0})
// 		}
// 	}

// 	return arr
// }
