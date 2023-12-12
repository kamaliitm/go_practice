package heaps

type UserBalance struct {
	userId  string
	balance float32
}

type MyMaxHeap struct {
	data []UserBalance
	size int
}

func NewMyMaxHeap() *MyMaxHeap {
	return &MyMaxHeap{
		data: []UserBalance{},
		size: 0,
	}
}

func (h *MyMaxHeap) add(userBalance UserBalance) {
	h.size += 1
	h.data = append(h.data, userBalance)
	h.heapifyUp(h.size - 1)
}

func (h *MyMaxHeap) poll() *UserBalance {
	if h.size <= 0 {
		return nil
	}

	userBalance := h.data[0]
	h.size -= 1
	h.data[0] = h.data[h.size]
	h.heapifyDown(0)

	return &userBalance
}

func (h *MyMaxHeap) peek() *UserBalance {
	if h.size <= 0 {
		return nil
	}

	return &h.data[0]
}

func (h *MyMaxHeap) heapifyUp(index int) {
	if h.size == 0 && !h.isValidIndex(index) {
		return
	}

	parentInd := parent(index)
	if parentInd >= 0 && h.data[index].balance > h.data[parentInd].balance {
		h.swap(parentInd, index)
		h.heapifyUp(parentInd)
	}
}

func (h *MyMaxHeap) heapifyDown(index int) {
	if h.size == 0 || !h.isValidIndex(index) {
		return
	}

	maxIndex := index
	maxBalance := h.data[index].balance
	if leftChildInd := leftChild(index); h.isValidIndex(leftChildInd) {
		if maxBalance < h.data[leftChildInd].balance {
			maxBalance = h.data[leftChildInd].balance
			maxIndex = leftChildInd
		}
	}

	if rightChildInd := rightChild(index); h.isValidIndex(rightChildInd) {
		if maxBalance < h.data[rightChildInd].balance {
			maxBalance = h.data[rightChildInd].balance
			maxIndex = rightChildInd
		}
	}

	if maxIndex != index {
		h.swap(index, maxIndex)
		h.heapifyDown(maxIndex)
	}

}

func (h *MyMaxHeap) isValidIndex(index int) bool {
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

func (h *MyMaxHeap) swap(i int, j int) {
	if h.size <= 0 || i <= 0 || j <= 0 || i >= h.size || j >= h.size {
		return
	}

	h.data[i], h.data[j] = h.data[j], h.data[i]
}
