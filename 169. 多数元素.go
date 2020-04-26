package main

func majorityElement(nums []int) int {
	cnt := 0
	ret := 0
	for _, num := range nums {
		if cnt == 0 {
			ret = num
		}

		if ret == num {
			cnt ++
		} else {
			cnt --
		}
	}
	return ret
}
