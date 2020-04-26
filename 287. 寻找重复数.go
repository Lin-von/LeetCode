package main

func findDuplicate(nums []int) int {
	p := nums[0]
	q := nums[0]

	p = nums[p]
	q = nums[nums[q]]

	for p != q {
		p = nums[p]
		q = nums[nums[q]]
	}

	s := nums[0]
	f := p
	for s != f {
		s = nums[s]
		f = nums[f]
	}

	return s
}
