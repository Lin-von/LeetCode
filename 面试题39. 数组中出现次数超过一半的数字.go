package main

func majorityElement(nums []int) int {
	tmp := 0
	cnt := 0

	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			tmp = nums[i]
			cnt ++
		} else {
			if nums[i] == tmp {
				cnt ++
			} else {
				cnt --
			}
		}

	}
	return tmp
}
