package main

func jump(nums []int) int {
	ret := 0
	c := 0
	f := 0
	for f < len(nums) - 1 {
		tmp := 0
		for i := c; i <= f; i ++ {
			if i + nums[i] > tmp {
				tmp = i + nums[i]
			}
		}
		c = f + 1
		f = tmp
		ret ++
	}
	return ret
}