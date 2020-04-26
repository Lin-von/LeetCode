package main

import "fmt"

var arr []int

func scan(cur []int, target int, index int, ret *[][]int)  {
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}

	if target < 0 || index == len(arr) {
		return
	}

	tmp := append(cur, arr[index])
	scan(tmp, target - arr[index], index, ret)
	scan(cur, target, index + 1, ret)
}

func combinationSum(candidates []int, target int) [][]int {
	arr = candidates
	c := []int{}
	ret := [][]int{}
	scan(c, target, 0, &ret)
	return ret
}