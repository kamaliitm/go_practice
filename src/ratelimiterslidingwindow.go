package main

// const CHANNEL_COUNT = 5

// type RateLimiter struct {
// 	serviceRequestLimits map[string]int32
// 	serviceRequestCounts map[string]map[int64]int32
// 	channels             []chan ChannelMessage
// }

// type ChannelMessage struct {
// 	serviceId          string
// 	timestampInSeconds int64
// 	chResponseOut      chan bool
// }

// // Returns whether a request should or shouldn't be accepted according to configured limits
// func (ratelimiter *RateLimiter) shouldAccept(serviceId string, timestampInSeconds int64) bool {
// 	if ratelimiter.channels == nil {
// 		return false
// 	}

// 	chIndex := int(hashString(serviceId)) % len(ratelimiter.channels)
// 	chOut := make(chan bool)
// 	go func() {
// 		ratelimiter.channels[chIndex] <- ChannelMessage{
// 			serviceId:          serviceId,
// 			timestampInSeconds: timestampInSeconds,
// 			chResponseOut:      chOut,
// 		}
// 	}()

// 	shouldAccept := <-chOut

// 	return shouldAccept
// }

// // Configure request per second limit for this service
// func (ratelimiter *RateLimiter) configure(serviceId string, requestsPerSecond int32) {
// 	if ratelimiter.serviceRequestLimits == nil {
// 		ratelimiter.serviceRequestLimits = map[string]int32{}
// 	}

// 	if ratelimiter.channels == nil {
// 		ratelimiter.setUpChannels()
// 	}

// 	ratelimiter.serviceRequestLimits[serviceId] = requestsPerSecond
// }

// func (ratelimiter *RateLimiter) setUpChannels() {
// 	ratelimiter.channels = []chan ChannelMessage{}
// 	for i := 0; i < CHANNEL_COUNT; i++ {
// 		ch := make(chan ChannelMessage)
// 		go listenRateLimit(ch, ratelimiter)
// 		ratelimiter.channels = append(ratelimiter.channels, ch)
// 	}
// }

// func listenRateLimit(ch chan ChannelMessage, ratelimiter *RateLimiter) {
// 	for {
// 		chMsg := <-ch
// 		shouldAccept := ratelimiter.shouldAccept0(chMsg.serviceId, chMsg.timestampInSeconds)
// 		chMsg.chResponseOut <- shouldAccept
// 	}
// }

// func (ratelimiter *RateLimiter) shouldAccept0(serviceId string, timestampInSeconds int64) bool {
// 	if ratelimiter.serviceRequestLimits == nil {
// 		return false
// 	}

// 	requestLimit, ok := ratelimiter.serviceRequestLimits[serviceId]
// 	if !ok {
// 		return false
// 	}

// 	if requestLimit <= 0 {
// 		return false
// 	}

// 	if ratelimiter.serviceRequestCounts == nil {
// 		ratelimiter.serviceRequestCounts = map[string]map[int64]int32{}
// 		ratelimiter.serviceRequestCounts[serviceId] = map[int64]int32{}
// 		ratelimiter.serviceRequestCounts[serviceId][timestampInSeconds] = 1
// 		return true
// 	}

// 	requestCountWindowMap, ok2 := ratelimiter.serviceRequestCounts[serviceId]
// 	if !ok2 {
// 		ratelimiter.serviceRequestCounts[serviceId] = map[int64]int32{}
// 		ratelimiter.serviceRequestCounts[serviceId][timestampInSeconds] = 1
// 		return true
// 	}

// 	requestCount, ok3 := requestCountWindowMap[timestampInSeconds]
// 	if !ok3 {
// 		ratelimiter.serviceRequestCounts[serviceId][timestampInSeconds] = 1
// 		return true
// 	}

// 	if requestCount < requestLimit {
// 		ratelimiter.serviceRequestCounts[serviceId][timestampInSeconds] += 1
// 		return true
// 	}

// 	return false

// }
