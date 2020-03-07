package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := 1; i <=amount; i++ {
		dp[i] = math.MaxInt32
		for _, j := range coins {
			if i - j >= 0 && dp[i - j] != math.MaxInt32 && dp[i - j] + 1 < dp[i] {
				dp[i] = dp[i - j] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[amount]
	}
}
