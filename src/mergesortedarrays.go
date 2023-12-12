package main

// func mergeSortedArrays(arrOfArrays [][]int) []int {
// 	res := []int{}

// 	noOfArrays := len(arrOfArrays)
// 	minHeap := &MinHeap{}
// 	currIndices := make([]int, noOfArrays)

// 	for i := 0; i < noOfArrays; i++ {
// 		currentArr := arrOfArrays[i]
// 		minHeap.insert(HeapElement{currentArr[0], i})
// 		currIndices[i] = 1
// 	}

// 	for {
// 		minElem := minHeap.extractMin()
// 		fmt.Println("heap after extracting", minElem.val, "heap", minHeap.data)
// 		res = append(res, minElem.val)
// 		fmt.Println("res", res)

// 		if len(minHeap.data) == 0 {
// 			break
// 		}

// 		minElemPickedFromArr := minElem.index
// 		currIndexToPick := currIndices[minElemPickedFromArr]
// 		arrToPickVal := arrOfArrays[minElemPickedFromArr]
// 		if currIndexToPick < len(arrToPickVal) {
// 			minHeap.insert(HeapElement{
// 				arrToPickVal[currIndexToPick],
// 				minElemPickedFromArr,
// 			})
// 			currIndices[minElemPickedFromArr] += 1
// 		}
// 	}

// 	return res
// }
