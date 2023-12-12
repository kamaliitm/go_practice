package main

import (
	"time"
)

type CacheV2 struct {
	data map[string]Value
	size int32
}

type Value struct {
	val      string
	lastUsed int64
}

// Get values from cache
// Return null if the value does not exist
func (cache *CacheV2) get(key string) string {
	if value, ok := cache.data[key]; ok {
		value.lastUsed = time.Now().UnixNano()
		return value.val
	}

	return "null"
}

// Set values in cache
func (cache *CacheV2) set(key string, value string) {
	if int32(len(cache.data)) == cache.size {
		// evict
		var evictKey string
		var minLastUsed int64 = -1
		for k, value := range cache.data {
			if minLastUsed == -1 {
				minLastUsed = value.lastUsed
				evictKey = k
			} else {
				if value.lastUsed < minLastUsed {
					evictKey = k
					minLastUsed = value.lastUsed
				}
			}
		}

		delete(cache.data, evictKey)
	}

	cache.data[key] = Value{
		val:      value,
		lastUsed: time.Now().UnixNano(),
	}

}

// Sets the size of cache.
// It is run once during creation of the cache
func (cache *CacheV2) configure(size int32) {
	cache.data = map[string]Value{}
	cache.size = size
}
