package main

func minIncrementForUnique(A []int) int {
	h := make([]int, 80000)
	for _, v := range A {
		h[v] ++
	}
	cnt := 0
	ret := 0
	l := len(A)
	for i := 0; i < 80000; i++ {
		if h[i] > 1 {
			h[i + 1] += h[i] - 1
			ret += h[i] - 1
			cnt ++
		}
		if cnt == l {
			break
		}
	}
	return ret
}
