package main

import "math"

var mmin int
var nums []int

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func bs(l, r int)  {
	if r - l == 1 {
		mmin = min(nums[l], nums[r])
		return
	}

	x := (l + r) / 2
	if nums[x] < nums[x - 1] && nums[x] < nums[x + 1] {
		mmin = nums[x]
		return
	} else if nums[x] > nums[r] {
		bs(x, r)
	} else if nums[x] > nums[l] {
		bs(l, x)
	} else {
		fl := false
		for i := l; i < x; i++ {
			if nums[i] > nums[i + 1] {
				fl = true
				break
			}
		}
		if fl {
			bs(l, x)
		} else {
			bs(x, r)
		}
	}
}

func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	l := 0
	r := len(numbers) - 1
	nums = numbers
	mmin = math.MaxInt32
	bs(l, r)
	return mmin
}
