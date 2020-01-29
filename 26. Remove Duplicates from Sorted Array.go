package main

// 双指针 swap
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	cur := 0
	pos := cur + 1

	for pos != n {
		for pos < n && nums[pos] == nums[cur] {
			pos ++
		}

		if pos == n {
			break
		}

		cur ++
		nums[cur] = nums[pos]
		pos ++
	}

	return cur + 1
}
