package main

import "math"

func maxProfit(stockPrices []int32) int32 {
	var maxProfit int32 = 0
	var minPriceSoFar int32 = 0

	for _, price := range stockPrices {
		if minPriceSoFar == 0 || minPriceSoFar > price {
			minPriceSoFar = price
		}

		maxProfit = int32(math.Max(float64(price-minPriceSoFar), float64(maxProfit)))
	}

	return maxProfit
}
