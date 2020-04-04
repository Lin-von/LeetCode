package main

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		s1 := strconv.Itoa(nums[i])
		s2 := strconv.Itoa(nums[j])
		return s1 + s2 < s2 + s1
	})
	ret := ""
	for _, num := range nums {
		ret += strconv.Itoa(num)
	}

	return ret
}


