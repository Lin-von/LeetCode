package main

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := 0
	r := len(nums) - 1
	if nums[l] < nums[r] {
		return nums[l]
	}
	for l < r {
		x := (l + r) / 2
		if nums[x] < nums[r] {
			r = x
		} else if nums[x] > nums[r] {
			l = x + 1
		} else {
			r --
		}
		if nums[l] < nums[r] {
			return nums[l]
		}
	}
	return nums[l]
}
