package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	p := 0
	q := 1
	for q < len(intervals) {
		if intervals[q][0] <= intervals[p][1] {
			if intervals[p][1] < intervals[q][1] {
				intervals[p][1] = intervals[q][1]
			}
			q ++
		} else {
			p ++
			intervals[p] = intervals[q]
			q ++
		}
	}

	return intervals[:p]
}
