package main

func exchange(nums []int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		for l < len(nums) && nums[l] & 1 == 1 {
			l ++
		}
		for r >= 0 && nums[r] & 1 == 0 {
			r --
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
		l ++
		r --
	}
	return nums
}
