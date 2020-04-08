package main

import (
	"sort"
)

var b []int

type Fenwick struct {
	val int
	idx int
}

func lowbit(x int) int {
	return x & -x
}

func update(i, val int) {
	for i < len(b) {
		b[i] += val
		i += lowbit(i)
	}
}

func get(i int) int {
	ret := 0
	for i > 0 {
		ret += b[i]
		i -= lowbit(i)
	}
	return ret
}

func reversePairs(nums []int) int {
	a := []Fenwick{}
	for i := 0; i < len(nums); i++ {
		a = append(a, Fenwick{nums[i], i + 1})
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].val == a[j].val {
			return a[i].idx < a[j].idx
		} else {
			return a[i].val < a[j].val
		}
	})

	b = make([]int, len(nums) + 1)

	ret := 0
	for i := 0; i < len(a); i++ {
		update(a[i].idx, 1)
		ret += a[i].idx - get(a[i].idx)
	}
	return ret
}
