package heaps

type MyMinHeap struct {
	data []UserBalance
	size int
}

func NewMyMinHeap() *MyMinHeap {
	return &MyMinHeap{
		data: []UserBalance{},
		size: 0,
	}
}

func (h *MyMinHeap) add(userBalance UserBalance) {
	h.size += 1
	h.data = append(h.data, userBalance)
	h.heapifyUp(h.size - 1)
}

func (h *MyMinHeap) poll() *UserBalance {
	if h.size <= 0 {
		return nil
	}

	userBalance := h.data[0]
	h.size -= 1
	h.data[0] = h.data[h.size]
	h.heapifyDown(0)

	return &userBalance
}

func (h *MyMinHeap) peek() *UserBalance {
	if h.size <= 0 {
		return nil
	}

	return &h.data[0]
}

func (h *MyMinHeap) heapifyUp(index int) {
	if h.size == 0 && !h.isValidIndex(index) {
		return
	}

	parentInd := parent(index)
	if parentInd >= 0 && h.data[index].balance < h.data[parentInd].balance {
		h.swap(parentInd, index)
		h.heapifyUp(parentInd)
	}
}

func (h *MyMinHeap) heapifyDown(index int) {
	if h.size == 0 || !h.isValidIndex(index) {
		return
	}

	minIndex := index
	minBalance := h.data[index].balance
	if leftChildInd := leftChild(index); h.isValidIndex(leftChildInd) {
		if minBalance > h.data[leftChildInd].balance {
			minBalance = h.data[leftChildInd].balance
			minIndex = leftChildInd
		}
	}

	if rightChildInd := rightChild(index); h.isValidIndex(rightChildInd) {
		if minBalance > h.data[rightChildInd].balance {
			// minBalance = h.data[rightChildInd].balance
			minIndex = rightChildInd
		}
	}

	if minIndex != index {
		h.swap(index, minIndex)
		h.heapifyDown(minIndex)
	}

}

func (h *MyMinHeap) isValidIndex(index int) bool {
	return index >= 0 && index < h.size
}

func (h *MyMinHeap) swap(i int, j int) {
	if h.size <= 0 || i <= 0 || j <= 0 || i >= h.size || j >= h.size {
		return
	}

	h.data[i], h.data[j] = h.data[j], h.data[i]
}
