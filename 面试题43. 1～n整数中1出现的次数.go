package main

func countDigitOne(n int) int {
	ret := 0
	for i := 1; i <= n; i *= 10 {
		a := n / i
		b := n % i
		ret += (a + 8) / 10 * i
		if a % 10 == 1 {
			ret += b + 1
		}
	}
	return ret
}
