package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}

	for i := 0; i <= len(nums) - 4; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		for j := i + 1; j <= len(nums) - 3; j++ {
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				tmp := nums[i] + nums[j] + nums[l] + nums[r]
				if tmp == target {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					l ++
					for l < r && nums[l] == nums[l - 1] {
						l ++
					}
					r --
					for l < r && nums[r] == nums[r + 1] {
						r --
					}
				} else if tmp < target {
					l ++
				} else {
					r --
				}
			}
		}
	}
	return ret
}
