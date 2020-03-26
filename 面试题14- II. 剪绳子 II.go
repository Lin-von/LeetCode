package main

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	ret := 1
	for n > 4 {
		n -= 3
		ret *= 3
		ret %= 1e9+7
	}
	return n * ret % (1e9+7)
}
