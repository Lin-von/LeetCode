package main

import "math"

func div(a, b int) int {
	if a > b {
		return 0
	}

	cnt := 1
	tmp := b
	for tmp + tmp > a && tmp + tmp < 0 {
		cnt += cnt
		tmp += tmp
	}
	return cnt + div(a - tmp, b)
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}

	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}

	if dividend > 0 {
		dividend = -dividend
	}

	if divisor > 0 {
		divisor = -divisor
	}

	ret := div(dividend, divisor)
	if sign == -1 {
		return -ret
	}
	return ret
}
