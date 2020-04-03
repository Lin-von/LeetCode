package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tmp := nums[0]
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > tmp + nums[i] {
			tmp = nums[i]
		} else {
			tmp += nums[i]
		}

		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}
