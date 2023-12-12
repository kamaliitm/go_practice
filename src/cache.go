// // The cache must accept GET and SET requests for different keys, and process these requests in concurrent threads
package main

// const OP_GET = "get"
// const OP_SET = "set"
// const NULL = "null"
// const DEFAULT_THRESHOLD = 500
// const DEFAULT_CHANNEL_COUNT = 5

// // evict algorithm will be a combination of LRU and LFU;
// // remove entries based on LRU, if there is a tie consider LFU.
// type Cache struct {
// 	data      map[string]string
// 	threshold uint
// 	channels  []chan Operation
// }

// type Operation struct {
// 	op    string
// 	key   string
// 	value string
// 	chOut chan string
// }

// func (cache *Cache) configureNewCache(threshold uint, noOfChannels uint) {
// 	cache.threshold = threshold
// 	cache.channels = []chan Operation{}
// 	for i := 0; i < int(noOfChannels); i++ {
// 		ch := make(chan Operation)
// 		go listen(ch, cache)
// 		cache.channels = append(cache.channels, ch)
// 	}
// }

// // retrieve this element from the cache
// func (cache *Cache) get(key string) string {

// 	if cache.channels == nil {
// 		return NULL
// 	}

// 	chIndex := getChannelIndex(key, len(cache.channels))
// 	ch := cache.channels[chIndex]
// 	chOut := make(chan string)
// 	ch <- Operation{op: OP_GET, key: key, chOut: chOut}

// 	return <-chOut
// }

// // set this element in the cache
// func (cache *Cache) set(key string, value string) {

// 	if cache.channels == nil {
// 		cache.configureNewCache(DEFAULT_THRESHOLD, DEFAULT_CHANNEL_COUNT)
// 	}

// 	chIndex := getChannelIndex(key, len(cache.channels))
// 	ch := cache.channels[chIndex]
// 	msg := Operation{op: OP_SET, key: key, value: value}

// 	ch <- msg
// }

// func getChannelIndex(key string, poolSize int) int {
// 	return int(hashString(key)) % poolSize

// }

// func listen(ch chan Operation, cache *Cache) {
// 	for {
// 		operation := <-ch
// 		if operation.op == OP_GET {
// 			val := cache.get0(operation.key)
// 			operation.chOut <- val
// 		} else {
// 			cache.set0(operation.key, operation.value)
// 		}
// 	}
// }

// func (cache *Cache) get0(key string) string {
// 	if val, ok := cache.data[key]; !ok {
// 		return NULL
// 	} else {
// 		return val
// 	}
// }

// func (cache *Cache) set0(key string, value string) {
// 	if cache.data == nil {
// 		cache.data = map[string]string{}
// 	}

// 	cache.data[key] = value
// 	//fmt.Println("data", cache.data)
// }
