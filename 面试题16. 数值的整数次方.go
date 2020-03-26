package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n *= -1
	}

	ret := 1.0
	tmp := x
	for n != 0 {
		if (n & 1 == 1) {
			ret *= tmp
		}
		tmp *= tmp
		n >>= 1
	}
	return ret
}
