package main

func missingNumber(nums []int) int {
	l := 0
	r := len(nums)
	x := (l + r) >> 1

	for l <= r {
		x = (l + r) >> 1
		if x == len(nums) {
			return x
		}
		if x == nums[x] {
			l = x + 1
		} else {
			if x == 0 || x - 1 == nums[x - 1] {
				return x
			}
			r = x - 1
		}
	}
	return x
}
