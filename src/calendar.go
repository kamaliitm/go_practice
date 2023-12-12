package main

import "sort"

type CalEvent struct {
	start int64
	end   int64
}

type EventTime struct {
	t       int64
	isStart bool
}

func findMaxSimulateneousEvents(events []CalEvent) int {
	eventTimeArr := []EventTime{}
	for _, e := range events {
		eventTimeArr = append(eventTimeArr, EventTime{e.start, true})
		eventTimeArr = append(eventTimeArr, EventTime{e.end, false})
	}

	sort.Slice(eventTimeArr, func(i, j int) bool {
		return eventTimeArr[i].t < eventTimeArr[j].t
	})

	maxSimultaneousEventCount := 1
	simultaneousEventCount := 0
	for _, et := range eventTimeArr {
		if et.isStart {
			simultaneousEventCount += 1
			if simultaneousEventCount > maxSimultaneousEventCount {
				maxSimultaneousEventCount = simultaneousEventCount
			}
		} else {
			simultaneousEventCount -= 1
		}
	}

	return maxSimultaneousEventCount
}
