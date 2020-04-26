package main


// 取 当前最大和+当前值与当前值中的较大值 维护即可
func maxSubArray(nums []int) int {
	n := len(nums)
	if (n == 0) {
		return 0
	}
	result := nums[0]
	for sum, i := nums[0], 1; i < n; i++ {
		if sum + nums[i] > nums[i] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > result {
			result = sum
		}
	}

	return result
}

