package main

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	dp := make([]int, num + 1)
	dp[0] = 0
	dp[1] = 1
	step := 2
	for i := 2; i <= num; i++ {
		if i == step * 2 {
			step *= 2
		}
		dp[i] = dp[i - step] + 1
	}


	return dp
}
