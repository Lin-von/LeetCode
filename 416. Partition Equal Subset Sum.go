package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum & 1 == 1 {
		return false
	}

	dp := make([]bool, sum / 2 + 1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum / 2; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j - nums[i]]
		}
	}
	return dp[sum / 2]
}
