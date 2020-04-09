package main

func singleNumbers(nums []int) []int {
	ret := make([]int, 2)
	tmp := 0
	for _, num := range nums {
		tmp ^= num
	}
	x := 1
	for x & tmp == 0 {
		x <<= 1
	}
	for _, num := range nums {
		if num & x == 0 {
			ret[0] ^= num
		} else {
			ret[1] ^= num
		}
	}
	return ret
}
