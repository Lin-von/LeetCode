package main

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
			continue
		}

		for i != nums[i] {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return 0
}
