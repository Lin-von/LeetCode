package main

func twoSum(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	for nums[l] + nums[r] != target {
		if nums[l] + nums[r] > target {
			r --
		} else {
			l ++
		}
	}
	return []int{nums[l], nums[r]}
}
