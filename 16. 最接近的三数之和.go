package main

import "sort"

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ret := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			if tmp == target {
				return target
			}
			if abs(tmp - target) < abs(ret - target) {
				ret = tmp
			}
			if target > tmp {
				l ++
			} else {
				r --
			}
		}
	}
	return ret
}
