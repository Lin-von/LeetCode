package main

func findTargetSumWays(nums []int, S int) int {
	if S > 1000 {
		return 0
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i ++ {
		dp[i] = make([]int, 2001)
	}

	dp[0][1000 - nums[0]] = 1
	dp[0][1000 + nums[0]] += 1
	for i := 1; i < len(nums); i ++ {
		for j := -1000; j <= 1000; j++ {
			if dp[i - 1][j + 1000] > 0 {
				dp[i][j - nums[i] + 1000] += dp[i - 1][j + 1000]
				dp[i][j + nums[i] + 1000] += dp[i - 1][j + 1000]
			}
		}
	}
	return dp[len(nums) - 1][S + 1000]
}
