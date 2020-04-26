package main

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	ret := 1
	for n > 4 {
		n -= 3
		ret *= 3
	}
	return n * ret
}
