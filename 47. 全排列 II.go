package main

import "sort"

var arr []int
var f []bool

func scan(cur []int, ret *[][]int)  {
	if len(cur) == len(arr) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}

	for i := 0; i < len(arr); i++ {
		if f[i] {
			continue
		}

		if i > 0 && !f[i - 1] && arr[i] == arr[i - 1] {
			continue
		}

		f[i] = true
		tmp := append(cur, arr[i])
		scan(tmp, ret)
		f[i] = false
	}

}

func permuteUnique(nums []int) [][]int {
	arr = nums
	c := []int{}
	ret := [][]int{}
	f = make([]bool, len(arr))
	sort.Ints(arr)
	scan(c, &ret)
	return ret
}