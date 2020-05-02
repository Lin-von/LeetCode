package main

import "sort"

var can []int
var tar int
var ret [][]int
var f []bool

func dfs(cur []int, pos, sum int)  {
	if sum == tar {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		ret = append(ret, tmp)
		return
	}

	if pos == len(can) || sum > tar {
		return
	}

	for i := pos; i < len(can); i++ {
		if f[i] {
			continue
		}

		if i > 0 && !f[i - 1] && can[i] == can[i - 1] {
			continue
		}

		f[i] = true
		cur = append(cur, can[i])
		dfs(cur, i + 1, sum + can[i])
		cur = cur[:len(cur) - 1]
		f[i] = false
	}

}

func combinationSum2(candidates []int, target int) [][]int {
	can = candidates
	tar = target
	ret = [][]int{}
	f = make([]bool, len(candidates))
	sort.Ints(can)
	dfs([]int{}, 0, 0)
	return ret
}
