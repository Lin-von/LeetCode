package main

import "math"

func findUnsortedSubarray(nums []int) int {
	min := math.MaxInt32
	max := math.MinInt32
	l , r := 0 , len(nums) - 1
	flag := false

	// 正向扫找到降序位置，找乱序区间内最小值
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i + 1] {
			flag = true
		}
		if flag && nums[i + 1] < min {
			min = nums[i + 1]
		}
	}

	flag = false
	// 逆向扫找到升序位置，找乱序区间内最大值
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] < nums[i - 1] {
			flag = true
		}
		if flag && nums[i - 1] > max {
			max = nums[i - 1]
		}
	}

	// 找到最小值的正确位置
	for l < len(nums) && nums[l] <= min {
		l ++
	}

	// 找到最大值的正确位置
	for r >= 0 && nums[r] >= max {
		r --
	}
	
	if l > r {
		return 0
	}

	return r - l + 1
}
