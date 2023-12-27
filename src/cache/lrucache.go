package cache

import "errors"

type LRUNode struct {
	key  string
	val  int
	next *LRUNode
	prev *LRUNode
}

type LRUCachePolicy struct {
	capacity int32
	data     map[string]*LRUNode
	head     *LRUNode
	tail     *LRUNode
	size     int32
}

func NewLRUCachePolicy(capacity int32) CachePolicy {
	return &LRUCachePolicy{
		capacity: capacity,
		data:     map[string]*LRUNode{},
	}
}

func (lru *LRUCachePolicy) set(key string, val int) {
	node, ok := lru.data[key]
	if ok {
		lru.removeNode(node)
		node.val = val
		lru.addNode(node)
		lru.data[key] = node
		return
	}

	if lru.size == lru.capacity {
		delete(lru.data, lru.head.key)
		lru.removeHead()
	}

	node = &LRUNode{key: key, val: val}
	lru.addNode(node)
	lru.data[key] = node
	lru.size += 1
}

func (lru *LRUCachePolicy) get(key string) (int, error) {
	if lru.size == 0 {
		return -1, errors.New("cache is empty")
	}

	node, ok := lru.data[key]
	if !ok {
		return -1, errors.New("key not found")
	}

	lru.removeNode(node)
	lru.addNode(node)

	return node.val, nil
}

func (lru *LRUCachePolicy) addNode(node *LRUNode) {
	if lru.head == nil {
		lru.head = node
		lru.tail = node
		return
	}

	lru.tail.next = node
	node.prev = lru.tail
	lru.tail = node
}

func (lru *LRUCachePolicy) removeNode(node *LRUNode) {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}
}

func (lru *LRUCachePolicy) removeHead() {
	if lru.head != nil {
		lru.head = lru.head.next
	}
}
