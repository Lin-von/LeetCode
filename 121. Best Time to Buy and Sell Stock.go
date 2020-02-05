package main

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt32
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i] - min > profit {
			profit = prices[i] - min
		}
	}
	return profit
}