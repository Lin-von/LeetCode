package main

func findContinuousSequence(target int) [][]int {
	ret := [][]int{}
	l := 1
	r := 2
	for l < r {
		x := ((l + r) * (r - l + 1) + 1) / 2
		if x == target {
			tmp := []int{}
			for i := l; i <= r; i++ {
				tmp = append(tmp, i)
			}
			ret = append(ret, tmp)
			l ++
		} else if x < target {
			r ++
		} else {
			l ++
		}
	}
	return ret
}
