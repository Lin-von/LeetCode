package main

func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	tmp := 1
	for i := 0; i < len(nums); i ++ {
		ret[i] = tmp
		tmp *= nums[i]
	}
	tmp = 1
	for i := len(nums) - 1; i >= 0; i -- {
		ret[i] *= tmp
		tmp *= nums[i]
	}
	return ret
}
