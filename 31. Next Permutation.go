package main

func nextPermutation(nums []int)  {
	var pos int
	for pos = len(nums) - 1; pos > 0 && nums[pos] <= nums[pos - 1]; pos-- {}
	if pos > 0 {
		for i := len(nums) - 1; i >= pos; i-- {
			if nums[i] > nums[pos - 1] {
				nums[i], nums[pos - 1] = nums[pos - 1], nums[i]
				break
			}
		}
	}

	for i, j := pos, len(nums) - 1; i <= j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
