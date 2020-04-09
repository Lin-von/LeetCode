package main

func singleNumber(nums []int) int {
	ret := 0
	x := 1
	for i := 0; i < 32; i++ {
		cnt := 0
		for _, num := range nums {
			if x & num == x {
				cnt ++
			}
		}
		if cnt % 3  == 1 {
			ret |= x
		}
		x <<= 1
	}
	return ret
}
