package medianfinder

type MinHeap struct {
	data []int
	size int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: []int{},
		size: 0,
	}
}

func (h *MinHeap) add(num int) {
	h.size += 1
	h.data = append(h.data, num)
	h.heapifyUp(h.size - 1)
}

func (h *MinHeap) poll() int {
	if h.size <= 0 {
		return -1
	}

	num := h.data[0]
	h.size -= 1
	h.data[0] = h.data[h.size]
	h.data = h.data[:h.size]
	h.heapifyDown(0)

	return num
}

func (h *MinHeap) peek() int {
	if h.size <= 0 {
		return -1
	}

	return h.data[0]
}

func (h *MinHeap) isEmpty() bool {
	return h.size <= 0
}

func (h *MinHeap) heapifyUp(index int) {
	if h.size == 0 && !h.isValidIndex(index) {
		return
	}

	parentInd := parent(index)
	if parentInd >= 0 && h.data[index] < h.data[parentInd] {
		h.swap(parentInd, index)
		h.heapifyUp(parentInd)
	}
}

func (h *MinHeap) heapifyDown(index int) {
	if h.size == 0 || !h.isValidIndex(index) {
		return
	}

	minIndex := index
	minNum := h.data[index]
	if leftChildInd := leftChild(index); h.isValidIndex(leftChildInd) {
		if minNum > h.data[leftChildInd] {
			minNum = h.data[leftChildInd]
			minIndex = leftChildInd
		}
	}

	if rightChildInd := rightChild(index); h.isValidIndex(rightChildInd) {
		if minNum > h.data[rightChildInd] {
			// minBalance = h.data[rightChildInd].balance
			minIndex = rightChildInd
		}
	}

	if minIndex != index {
		h.swap(index, minIndex)
		h.heapifyDown(minIndex)
	}

}

func (h *MinHeap) isValidIndex(index int) bool {
	return index >= 0 && index < h.size
}

func (h *MinHeap) swap(i int, j int) {
	if h.size <= 0 || i < 0 || j < 0 || i >= h.size || j >= h.size {
		return
	}

	h.data[i], h.data[j] = h.data[j], h.data[i]
}
