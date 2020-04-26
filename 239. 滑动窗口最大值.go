package main

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	if k == 1 {
		return nums
	}
	l := make([]int, len(nums))
	r := make([]int, len(nums))
	l[0] = nums[0]
	r[len(nums) - 1] = nums[len(nums) - 1]
	for i := 1; i < len(nums); i++ {
		if i % k == 0 {
			l[i] = nums[i]
		} else {
			if nums[i] > l[i - 1] {
				l[i] = nums[i]
			} else {
				l[i] = l[i - 1]
			}
		}

		tmp := len(nums) - 1 - i
		if (tmp + 1) % k == 0 {
			r[tmp] = nums[tmp]
		} else {
			if nums[tmp] > r[tmp + 1] {
				r[tmp] = nums[tmp]
			} else {
				r[tmp] = r[tmp + 1]
			}
		}
	}

	ret := []int{}
	for i := 0; i < len(nums) - k + 1; i++ {
		if r[i] > l[i + k - 1] {
			ret = append(ret, r[i])
		} else {
			ret = append(ret, l[i + k - 1])
		}
	}
	return ret
}
