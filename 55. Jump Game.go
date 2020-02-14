package main

func canJump(nums []int) bool {
	pos := 0
	for pos < len(nums) {
		i := pos
		for ; i <= pos + nums[pos]; i++ {
			if i >= len(nums) - 1 || i + nums[i] >= len(nums) - 1 {
				pos = len(nums)
				break
			}

			if i + nums[i] > pos + nums[pos] {
				pos = i
				break
			}
		}
		if pos != len(nums) && i == pos + nums[pos] + 1 {
			break
		}
	}

	if pos >= len(nums) {
		return true
	} else {
		return false
	}
}
