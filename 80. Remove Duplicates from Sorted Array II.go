package main

func removeDuplicates(nums []int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		if i < 2 || nums[i] != nums[ret - 2] {
			nums[ret] = nums[i]
			ret ++
		}
	}
	return ret
}
