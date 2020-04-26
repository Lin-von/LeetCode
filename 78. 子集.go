package main

// n位的数组有2^n个排列，用二进制方式即可表示出所有的排列方式，取对应的位置的元素即可
func subsets(nums []int) [][]int {
	ret := [][]int{}
	for i := 0; i < (1 << len(nums)); i++ {
		tmp := []int{}
		for j := 0; j < len(nums); j++ {
			if (i >> j) & 1 == 1 {
				tmp = append(tmp, nums[j])
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}
