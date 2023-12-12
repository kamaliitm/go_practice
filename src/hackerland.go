package main

import "sort"

func hackerlandRadioTransmitters(x []int32, k int32) int32 {

	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	numTrasmitters := int32(0)
	coveredDistance := int32(0)

	for i, house := range x {
		if house <= coveredDistance {
			continue
		}

		// find the furthest tower for this house for setting up transmitter
		j := i + 1
		for j < len(x) && (house+k) >= x[j] {
			j += 1
		}

		// set up transmitter at j-1
		numTrasmitters += 1
		coveredDistance = x[j-1] + k
	}

	return numTrasmitters
}
