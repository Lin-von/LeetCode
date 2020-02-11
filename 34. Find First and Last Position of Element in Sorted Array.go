package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	ret := []int{}
	l, r := 0, len(nums) - 1
	for l <= r {
		x := (l + r) / 2
		if nums[x] == target && (x == 0 || nums[x - 1] != target) {
			ret = append(ret, x)
			break
		}

		if nums[x] < target {
			l = x + 1
		} else {
			r = x - 1
		}
	}

	l, r = 0, len(nums) - 1
	for l <= r {
		x := (l + r) / 2
		if nums[x] == target && (x == len(nums) - 1 || nums[x + 1] != target) {
			ret = append(ret, x)
			break
		}

		if nums[x] > target {
			r = x - 1
		} else {
			l = x + 1
		}
	}

	if len(ret) == 0 {
		ret = []int{-1, -1}
	}
	return ret
}
