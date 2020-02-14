package main

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := [][]int{}
	p := 0
	for p < len(intervals) && intervals[p][1] <= newInterval[0] {
		ret = append(ret, intervals[p])
		p ++
	}

	if len(ret) > 0 && ret[len(ret) - 1][1] == newInterval[0] {
		ret[len(ret) - 1][1] = newInterval[1]
	}  else if p < len(intervals) && intervals[p][0] < newInterval[0] {
		ret = append(ret, intervals[p])
		if newInterval[1] > ret[len(ret) - 1][1] {
			ret[len(ret) - 1][1] = newInterval[1]
		}
		p ++
	} else {
		ret = append(ret, newInterval)
	}

	for p < len(intervals) && intervals[p][1] <= ret[len(ret) - 1][1] {
		p ++
	}

	if p < len(intervals) && intervals[p][0] <= ret[len(ret) - 1][1] {
		ret[len(ret) - 1][1] = intervals[p][1]
		p++
	}

	for p < len(intervals) {
		ret = append(ret, intervals[p])
		p ++
	}

	return ret

}
