package main

const DEFAULT_VAL = -1

type MinHeap struct {
	data     []int32
	heapSize int
}

func (h *MinHeap) peek() int32 {
	if h.data == nil || h.heapSize == 0 {
		return DEFAULT_VAL
	}

	return h.data[0]
}

func (h *MinHeap) extractMin() int32 {
	if h.data == nil || len(h.data) == 0 {
		return DEFAULT_VAL
	}

	min := h.data[0]

	h.heapSize -= 1
	h.data[0] = h.data[h.heapSize]
	h.data = h.data[:h.heapSize]

	h.heapifyDown(0)

	return min
}

func (h *MinHeap) add(key int32) {
	if h.data == nil {
		h.data = []int32{}
		h.heapSize = 0
	}

	h.data = append(h.data, key)
	h.heapSize += 1
	h.heapifyUp(h.heapSize - 1)
}

func (h *MinHeap) heapifyDown(index int) {
	if index > h.heapSize-1 {
		return
	}
	minVal := h.data[index]
	minIndex := index

	leftChildIndex := getLeftChildIndex(index)
	if leftChildIndex < h.heapSize && h.data[leftChildIndex] < minVal {
		minVal = h.data[leftChildIndex]
		minIndex = leftChildIndex
	}

	rightChildIndex := getRightChildIndex(index)
	if rightChildIndex < h.heapSize && h.data[rightChildIndex] < minVal {
		minVal = h.data[rightChildIndex]
		minIndex = rightChildIndex
	}

	if minIndex != index {
		h.swap(index, minIndex)
		h.heapifyDown(minIndex)
	}
}

func (h *MinHeap) heapifyUp(index int) {
	parentIndex := getParentIndex(index)

	if parentIndex >= 0 {
		if h.data[parentIndex] > h.data[index] {
			h.swap(parentIndex, index)
			h.heapifyUp(parentIndex)
		}
	}
}

func (h *MinHeap) swap(i int, j int) {
	if i < h.heapSize && j < h.heapSize {
		temp := h.data[i]
		h.data[i] = h.data[j]
		h.data[j] = temp
	}
}

func getLeftChildIndex(index int) int {
	return 2*index + 1
}

func getRightChildIndex(index int) int {
	return 2*index + 2
}

func getParentIndex(index int) int {
	if index == 0 {
		return DEFAULT_VAL
	}
	return (index - 1) / 2
}

func cookies(k int32, A []int32) int32 {
	// Write your code here
	minHeap := &MinHeap{}
	for _, key := range A {
		minHeap.add(key)
	}

	// fmt.Println("minheap", minHeap, "k", k)

	noOfIterations := 0

	for minHeap.heapSize > 1 && minHeap.peek() < k {
		leastSweet := minHeap.extractMin()
		secondLeastSweet := minHeap.extractMin()
		// fmt.Println("secondLeastSweet", secondLeastSweet, "minheap", minHeap)

		newSweet := leastSweet + (2 * secondLeastSweet)
		// fmt.Println("newSweet", newSweet, "minheap", minHeap)
		minHeap.add(newSweet)

		noOfIterations += 1
		// fmt.Println("noOfIterations", noOfIterations, "minheap", minHeap)
	}

	if minHeap.peek() >= k {
		return int32(noOfIterations)
	}

	return DEFAULT_VAL

	// for {
	// 	leastSweet := minHeap.extractMin()
	// 	// fmt.Println("leastSweet", leastSweet, "minheap", minHeap)
	// 	if leastSweet >= k {
	// 		break
	// 	}

	// 	if minHeap.heapSize == 0 {
	// 		return DEFAULT_VAL
	// 	}

	// 	secondLeastSweet := minHeap.extractMin()
	// 	// fmt.Println("secondLeastSweet", secondLeastSweet, "minheap", minHeap)

	// 	newSweet := leastSweet + (2 * secondLeastSweet)
	// 	// fmt.Println("newSweet", newSweet, "minheap", minHeap)
	// 	minHeap.add(newSweet)

	// 	noOfIterations += 1
	// 	// fmt.Println("noOfIterations", noOfIterations, "minheap", minHeap)
	// }

	// return int32(noOfIterations)
}
