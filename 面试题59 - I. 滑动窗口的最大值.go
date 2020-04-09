package main

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	if k == 1 {
		return nums
	}

	ret := []int{}
	s := []int{}
	for i := 0; i < len(nums); i ++ {
		if len(s) > 0 {
			if s[0] == i - k {
				s = s[1:]
			}
			for len(s) > 0 && nums[s[len(s) - 1]] < nums[i] {
				s = s[:len(s) - 1]
			}
		}
		s = append(s, i)
		if i < k - 1 {
			continue
		}
		ret = append(ret, nums[s[0]])
	}
	return ret
}
