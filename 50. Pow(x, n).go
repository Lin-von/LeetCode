package main

func myPow(x float64, n int) float64 {
	ret := 1.0
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for n != 0 {
		if n & 1 == 1 {
			ret *= x
		}
		x *= x
		n >>= 1
	}
	return ret
}
