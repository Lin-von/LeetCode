package main

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		x := (l + r) / 2
		if nums[x] == target {
			return x
		} else {

			if nums[l] <= nums[x] {
				if target >= nums[l] && target < nums[x] {
					r = x - 1
				} else {
					l = x + 1
				}
			} else {
				if target <= nums[r] && target > nums[x] {
					l = x + 1
				} else {
					r = x - 1
				}
			}

		}

	}

	return -1
}
