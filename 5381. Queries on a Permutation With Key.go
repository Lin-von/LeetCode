package main

func processQueries(queries []int, m int) []int {
	p := make([]int, m + 1)
	for i := 1; i < len(p); i++ {
		p[i] = i - 1
	}

	ret := []int{}
	for _, q := range queries {
		t := p[q]
		ret = append(ret, t)
		for i := 0; i < len(p); i++ {
			if p[i] < t {
				p[i] ++
			}
		}
		p[q] = 0
	}
	return ret
}
