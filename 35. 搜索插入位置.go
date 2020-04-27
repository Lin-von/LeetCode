package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		x := (l + r) >> 1
		if nums[x] == target {
			return x
		} else if nums[x] < target {
			l = x + 1
		} else {
			r = x - 1
		}
	}
	return l
}
