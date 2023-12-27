package cache

type LFUMinHeapNode struct {
	key         string
	val         int
	accessCount int
}

type LFUMinHeap struct {
	data []*LFUMinHeapNode
	size int32
}

func NewLFUMinHeap() *LFUMinHeap {
	return &LFUMinHeap{
		data: []*LFUMinHeapNode{},
	}
}

// return index in the data list
func (h *LFUMinHeap) add(val int) int {
	return 0
}

type LFUCachePolicy struct {
	capacity int32
	data     map[string]*LFUMinHeapNode
	size     int32
	minheap  *LFUMinHeap
}

func NewLFUCachePolicy(capacity int32) CachePolicy {
	return &LFUCachePolicy{
		capacity: capacity,
		data:     map[string]*LFUMinHeapNode{},
		minheap:  nil,
	}
}

func (c *LFUCachePolicy) set(key string, val int) {

}

func (c *LFUCachePolicy) get(key string) (int, error) {
	return -1, nil
}
