package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}

	r := len(nums) - 1
	for i := r - 2; i >= 0; i-- {
		for j := i + 2; j <= r; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k] + nums[i] * nums[k] * nums[j] + dp[k][j])
			}
		}
	}
	return dp[0][r]
}
