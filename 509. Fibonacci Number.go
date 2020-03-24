package main

func fib(N int) int {
	if N < 2 {
		return N
	}
	p := 0
	q := 1
	ret := 0
	for i := 2; i <= N; i++ {
		ret = (p + q) % (1e9 + 7)
		p = q
		q = ret
	}
	return ret
}