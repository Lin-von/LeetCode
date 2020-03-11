package main

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] > people[j][0]
		}
		return people[i][1] < people[j][1]
	})
	ret := [][]int{}
	for _, each := range people {
		ret = append(ret, each)
		copy(ret[each[1] + 1:], ret[each[1]:])
		ret[each[1]] = each
	}
	return ret
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0]
	})
	ret := make([][]int, len(people))
	for _, each := range people {
		tmp := each[1]
		pos := 0
		for {
			if len(ret[pos]) == 0 || ret[pos][0] == each[0] {
				if tmp == 0 {
					break
				}
				tmp --
			}

			pos ++
		}

		ret[pos] = each
	}
	return ret
}