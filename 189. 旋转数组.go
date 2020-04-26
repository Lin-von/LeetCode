package main

func rotate(nums []int, k int)  {
	k %= len(nums)
	if k == 0 {
		return
	}

	pos := len(nums) - k
	l := 0
	r := pos - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l ++
		r --
	}

	l = pos
	r = len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l ++
		r --
	}

	l = 0
	r = len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l ++
		r --
	}

}
