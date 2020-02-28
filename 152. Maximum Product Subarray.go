package main

func maxProduct(nums []int) int {
	maxSum := 1
	minSum := 1
	ret := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			maxSum, minSum = minSum, maxSum
		}

		if maxSum * nums[i] > nums[i] {
			maxSum *= nums[i]
		} else {
			maxSum = nums[i]
		}

		if minSum * nums[i] < nums[i] {
			minSum *= nums[i]
		} else {
			minSum = nums[i]
		}

		if maxSum > ret {
			ret = maxSum
		}
	}
	return ret
}
