package main

import "math"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	dp := make([]int, n + 1)
	dp[n] = 0
	t := []int{1, 7, 30}

	for i := n - 1; i >=0; i-- {
		dp[i] = math.MaxInt32
		k := i + 1
		for j := 0; j < 3; j++ {
			for k < n && days[k] < days[i] + t[j] {
				k ++
			}
			dp[i] = min(dp[i], dp[k] + costs[j])
		}
	}

	return dp[0]
}
