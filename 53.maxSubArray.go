package main



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

