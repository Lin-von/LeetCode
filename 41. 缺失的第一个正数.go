package main


func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i + 1 || nums[i] == 0 {
			continue
		}

		tmp := nums[i]
		pos := i
		for tmp > 0 && tmp <= len(nums) && tmp != pos + 1{
			pos = tmp - 1
			tmp = nums[pos]
			nums[pos] = pos + 1
		}
	}

	ret := 0
	for ret < len(nums) && nums[ret] == ret + 1 {
		ret ++
	}
	return ret + 1
}
