package main

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	p := 1
	q := 1
	ret := 0
	for i := 2; i <= n; i++ {
		ret = (p + q) % (1e9 + 7)
		p = q
		q = ret
	}
	return ret
}
