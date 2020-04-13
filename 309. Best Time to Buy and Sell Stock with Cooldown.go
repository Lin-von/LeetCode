package main

import "math"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxProfit(prices []int) int {
	dp0 := 0
	dp1 := math.MinInt32
	dppre0 := 0

	for i := 0; i < len(prices); i++ {
		tmp := dp0
		dp0 = max(dp0, dp1 + prices[i])
		dp1 = max(dp1, dppre0 - prices[i])
		dppre0 = tmp
	}
	return dp0
}
