package main

func sortColors(nums []int)  {
	p := 0
	q := len(nums) - 1

	if q <= 0 {
		return
	}

	for p < len(nums) && nums[p] == 0 {
		p ++
	}

	for q >= 0 && nums[q] == 2 {
		q --
	}

	for i := p; i <= q; i ++ {
		for (i >= p && i <= q) && (nums[i] == 0 || nums[i] == 2) {
			if nums[i] == 0 {
				nums[p], nums[i] = nums[i], nums[p]
				p ++
			} else if nums[i] == 2 {
				nums[q], nums[i] = nums[i], nums[q]
				q --
			}
		}
	}
}
