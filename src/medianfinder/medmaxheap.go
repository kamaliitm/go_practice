package medianfinder

type MaxHeap struct {
	data []int
	size int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: []int{},
		size: 0,
	}
}

func (h *MaxHeap) add(num int) {
	h.size += 1
	h.data = append(h.data, num)
	h.heapifyUp(h.size - 1)
}

func (h *MaxHeap) poll() int {
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

func (h *MaxHeap) peek() int {
	if h.size <= 0 {
		return -1
	}

	return h.data[0]
}

func (h *MaxHeap) isEmpty() bool {
	return h.size <= 0
}

func (h *MaxHeap) heapifyUp(index int) {
	if h.size == 0 && !h.isValidIndex(index) {
		return
	}

	parentInd := parent(index)
	if parentInd >= 0 && h.data[index] > h.data[parentInd] {
		h.swap(parentInd, index)
		h.heapifyUp(parentInd)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	if h.size == 0 || !h.isValidIndex(index) {
		return
	}

	maxIndex := index
	maxNum := h.data[index]
	if leftChildInd := leftChild(index); h.isValidIndex(leftChildInd) {
		if maxNum < h.data[leftChildInd] {
			maxNum = h.data[leftChildInd]
			maxIndex = leftChildInd
		}
	}

	if rightChildInd := rightChild(index); h.isValidIndex(rightChildInd) {
		if maxNum < h.data[rightChildInd] {
			maxNum = h.data[rightChildInd]
			maxIndex = rightChildInd
		}
	}

	if maxIndex != index {
		h.swap(index, maxIndex)
		h.heapifyDown(maxIndex)
	}

}

func (h *MaxHeap) isValidIndex(index int) bool {
	return index >= 0 && index < h.size
}

func parent(index int) int {
	if index <= 0 {
		return -1
	}

	return (index - 1) / 2
}

func leftChild(index int) int {
	return (2 * index) + 1
}

func rightChild(index int) int {
	return (2 * index) + 2
}

func (h *MaxHeap) swap(i int, j int) {
	if h.size <= 0 || i < 0 || j < 0 || i >= h.size || j >= h.size {
		return
	}

	h.data[i], h.data[j] = h.data[j], h.data[i]
}
