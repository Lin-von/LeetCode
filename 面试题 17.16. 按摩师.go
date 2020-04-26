package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func massage(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp[0] = nums[0]
	if nums[1] > nums[0] {
		dp[1] = nums[1]
	} else {
		dp[1] = dp[0]
	}
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i - 2] + nums[i], dp[i - 1])
	}
	return dp[len(nums) - 1]
}
