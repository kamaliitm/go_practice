package medianfinder

type MedianFinder struct {
	minHeap *MinHeap // bigger values
	maxHeap *MaxHeap // smaller values
}

func MFConstructor() MedianFinder {
	return MedianFinder{
		minHeap: NewMinHeap(),
		maxHeap: NewMaxHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.isEmpty() {
		this.minHeap.add(num)
	} else {
		if num >= this.minHeap.peek() {
			this.minHeap.add(num)
		} else {
			this.maxHeap.add(num)
		}
	}

	if this.minHeap.size > this.maxHeap.size+1 {
		this.maxHeap.add(this.minHeap.poll())
	} else if this.maxHeap.size > this.minHeap.size {
		this.minHeap.add(this.maxHeap.poll())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.size == this.maxHeap.size {
		return 0.5 * float64(this.minHeap.peek()+this.maxHeap.peek())
	}

	return float64(this.minHeap.peek())
}
