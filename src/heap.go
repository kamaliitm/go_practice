package main

// type HeapElement struct {
// 	val   int
// 	index int
// }

// type MinHeap struct {
// 	data []HeapElement
// }

// func (minHeap *MinHeap) insert(elem HeapElement) {
// 	if len(minHeap.data) == 0 {
// 		minHeap.data = append(minHeap.data, elem)
// 		fmt.Println("initialised heap with elem: ", elem, "and heap is", minHeap.data)
// 		return
// 	}

// 	minHeap.data = append(minHeap.data, elem)
// 	i := len(minHeap.data) - 1
// 	for i > 0 {
// 		parentIndex := (i - 1) / 2
// 		parent := minHeap.data[parentIndex]
// 		if parent.val > elem.val {
// 			minHeap.data[parentIndex] = elem
// 			minHeap.data[i] = parent
// 		}

// 		i = parentIndex
// 	}

// 	fmt.Println("heap after inserting elem: ", elem, " is ", minHeap.data)
// }

// func (minHeap *MinHeap) extractMin() *HeapElement {
// 	heapSize := len(minHeap.data)

// 	if heapSize == 0 {
// 		return nil
// 	}

// 	minVal := minHeap.data[0]
// 	minHeap.data[0] = minHeap.data[heapSize-1]
// 	minHeap.data = minHeap.data[:heapSize-1]
// 	minHeap.heapify(0)

// 	return &minVal
// }

// func (minHeap *MinHeap) heapify(i int) {
// 	heapSize := len(minHeap.data)

// 	if heapSize == 0 {
// 		return
// 	}

// 	for i < heapSize-1 {
// 		minElem := minHeap.data[i]
// 		minIndex := i

// 		leftChildIndex := (2 * i) + 1
// 		if leftChildIndex < heapSize {
// 			leftChild := minHeap.data[leftChildIndex]
// 			if leftChild.val < minElem.val {
// 				minElem = leftChild
// 				minIndex = leftChildIndex
// 			}
// 		}

// 		rightChildIndex := (2 * i) + 2
// 		if rightChildIndex < heapSize {
// 			rightChild := minHeap.data[rightChildIndex]
// 			if rightChild.val < minElem.val {
// 				minElem = rightChild
// 				minIndex = rightChildIndex
// 			}
// 		}

// 		if minIndex == i {
// 			break
// 		} else {
// 			// swap
// 			temp := minHeap.data[i]
// 			minHeap.data[i] = minElem
// 			minHeap.data[minIndex] = temp
// 			i = minIndex
// 		}
// 	}
// }
