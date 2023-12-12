package main

import "sort"

type Interval struct {
	left  EndPoint
	right EndPoint
}

type EndPoint struct {
	val      int32
	isClosed bool
}

func unionOfIntervals(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].left.val != intervals[j].left.val {
			return intervals[i].left.val < intervals[j].left.val
		}

		return intervals[i].left.isClosed && intervals[j].left.isClosed
	})

	res := []Interval{}
	resLen := 0
	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
			resLen += 1
			continue
		}

		lastInterval := &res[resLen-1]
		if (lastInterval.right.val > interval.left.val) ||
			(lastInterval.right.val == interval.left.val && (lastInterval.right.isClosed || interval.left.isClosed)) {
			if (interval.right.val > lastInterval.right.val) ||
				(interval.right.val == lastInterval.right.val && interval.right.isClosed) {
				// extend the existing interval
				lastInterval.right = interval.right
			}
		} else {
			// add this interval as a new one in the res
			res = append(res, interval)
			resLen += 1
		}
	}

	return res
}
