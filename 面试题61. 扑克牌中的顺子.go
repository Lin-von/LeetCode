package main

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)
	i := 0
	cnt := 0
	for nums[i] == 0 {
		cnt ++
		i ++
	}
	for i ++; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = nums[i - 1] + 1
			continue
		}

		if nums[i] == nums[i - 1] {
			return false
		}

		if nums[i] != nums[i - 1] + 1 {
			if nums[i] - nums[i - 1] - 1 <= cnt {
				cnt -= nums[i] - nums[i - 1] - 1
			} else {
				return false
			}
		}
	}
	return true
}
