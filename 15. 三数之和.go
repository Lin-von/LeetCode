package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			tmp :=  nums[i] + nums[l] + nums[r]
			if tmp == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l - 1] {
					l++
				}

				r--
				for r > l && nums[r] == nums[r + 1] {
					r--
				}
			} else if tmp < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return ret
}

