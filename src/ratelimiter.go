package main

import (
	"errors"
	"time"
)

type RateLimiterAlgorithm interface {
	shouldAccept(userId string, requestId string, timestampInSeconds int64) bool
}

type FixedWindow struct {
	currTimeStampInSeconds int64
	requestCountForUser    int32
}

type FixedWindowAlgorithm struct {
	timeWindow              int32
	requestLimitPerUser     int32
	requestCountMapForUsers map[string]FixedWindow
}

func NewFixedWindowAlgorithm(timeWindow int32, requestLimitPerUser int32) RateLimiterAlgorithm {
	return &FixedWindowAlgorithm{
		timeWindow:              timeWindow,
		requestLimitPerUser:     requestLimitPerUser,
		requestCountMapForUsers: map[string]FixedWindow{},
	}
}

func (fa *FixedWindowAlgorithm) shouldAccept(userId string, requestId string, timestampInSeconds int64) bool {
	currentWindow, ok := fa.requestCountMapForUsers[userId]
	if !ok || currentWindow.currTimeStampInSeconds < timestampInSeconds-int64(fa.timeWindow) {
		fa.requestCountMapForUsers[userId] = FixedWindow{
			currTimeStampInSeconds: timestampInSeconds,
			requestCountForUser:    1,
		}
		return true
	}

	if currentWindow.requestCountForUser == fa.requestLimitPerUser {
		return false
	}

	currentWindow.requestCountForUser += 1
	fa.requestCountMapForUsers[userId] = currentWindow

	return true
}

type SlidingWindowAlgorithm struct {
	timeWindow                int32
	requestLimitPerUser       int32
	requestTimestampsForUsers map[string][]int64
}

func NewSlidingWindowAlgorithm(timeWindow int32, requestLimitPerUser int32) RateLimiterAlgorithm {
	return &SlidingWindowAlgorithm{
		timeWindow:                timeWindow,
		requestLimitPerUser:       requestLimitPerUser,
		requestTimestampsForUsers: map[string][]int64{},
	}
}

func (sa *SlidingWindowAlgorithm) shouldAccept(userId string, requestId string, timestampInSeconds int64) bool {
	timestamps, ok := sa.requestTimestampsForUsers[userId]
	if !ok {
		sa.requestTimestampsForUsers[userId] = []int64{timestampInSeconds}
		return true
	}

	timestamps = filterTimestamps(timestamps, timestampInSeconds-int64(sa.timeWindow))
	if len(timestamps) == int(sa.requestLimitPerUser) {
		return false
	}

	timestamps = append(timestamps, timestampInSeconds)
	sa.requestTimestampsForUsers[userId] = timestamps

	return true
}

func filterTimestamps(timestamps []int64, minTime int64) (ret []int64) {
	for _, timestamp := range timestamps {
		if timestamp >= minTime {
			ret = append(ret, timestamp)
		}
	}

	return
}

type TimerWheelSlotInfo struct {
	currTimestampInSeconds int64
	requestCount           int32
}

type TimerWheelAlgorithm struct {
	timeout                    int32
	requestLimitPerUserPerSlot int32
	slotInfoForUsers           map[string]map[int32]TimerWheelSlotInfo
}

func NewTimerWheelAlgorithm(timeout int32, requestLimitPerUserPerSlot int32) RateLimiterAlgorithm {
	return &TimerWheelAlgorithm{
		timeout:                    timeout,
		requestLimitPerUserPerSlot: requestLimitPerUserPerSlot,
		slotInfoForUsers:           map[string]map[int32]TimerWheelSlotInfo{},
	}
}

func (twa *TimerWheelAlgorithm) shouldAccept(userId string, requestId string, timestampInSeconds int64) bool {
	slot := int32(timestampInSeconds % int64(twa.timeout))
	slotsInfo, ok := twa.slotInfoForUsers[userId]
	if !ok {
		slotsInfo = map[int32]TimerWheelSlotInfo{}
	}

	slotInfo, ok2 := slotsInfo[slot]
	if !ok2 || slotInfo.currTimestampInSeconds != timestampInSeconds {
		twa.slotInfoForUsers[userId][slot] = TimerWheelSlotInfo{
			currTimestampInSeconds: timestampInSeconds,
			requestCount:           1,
		}
		return true
	}

	if slotInfo.requestCount == twa.requestLimitPerUserPerSlot {
		return false
	}

	slotInfo.requestCount += 1
	twa.slotInfoForUsers[userId][slot] = slotInfo

	return true

}

type RateLimiter struct {
	algorithm RateLimiterAlgorithm
}

func NewRateLimiter(algo RateLimiterAlgorithm) RateLimiter {
	return RateLimiter{
		algorithm: algo,
	}
}

func (r *RateLimiter) canAcceptRequest(userId string, requestId string) (bool, error) {

	if r.algorithm.shouldAccept(userId, requestId, time.Now().Unix()) {
		return true, nil
	}

	return false, errors.New("too many requests from the user")
}

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
