package main

import "math"

type qnode struct {
	index int
	val int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func constrainedSubsetSum(nums []int, k int) int {
	ret := math.MinInt32
	dp := []qnode{}

	for i := 0; i < len(nums); i++ {
		if len(dp) == 0 {
			dp = append(dp, qnode{i, nums[i]})
			if nums[i] > ret {
				ret = nums[i]
			}
		} else {
			if dp[0].index < i - k {
				dp = dp[1:]
			}
			tmp := qnode{i, max(dp[0].val + nums[i], nums[i])}
			for len(dp) > 0 && dp[len(dp) - 1].val < tmp.val {
				dp = dp[:len(dp) - 1]
			}
			dp = append(dp, tmp)
			if tmp.val > ret {
				ret = tmp.val
			}
		}
	}
	return ret
}
