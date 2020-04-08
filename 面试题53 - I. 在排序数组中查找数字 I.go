package main


func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	var x int

	if len(nums) == 0 {
		return 0
	}

	for l <= r {
		x = (l + r) >> 1
		if nums[x] == target {
			break
		} else if nums[x] > target {
			r = x - 1
		} else {
			l = x + 1
		}
	}


	if nums[x] != target {
		return 0
	} else {
		cnt := 1
		for i := x - 1; i >= 0; i-- {
			if nums[i] != target {
				break
			}
			cnt ++
		}
		for i := x + 1; i < len(nums); i++ {
			if nums[i] != target {
				break
			}
			cnt ++
		}
		return cnt
	}
}
