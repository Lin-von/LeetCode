package main

func reverse(x int) int {
	result := 0
	for ; x != 0; {
		tmp := x % 10
		result = result * 10 + tmp
		x /= 10
	}
	if (result > 2147483648 - 1) || (result < -2147483648) {
		return 0
	}
		
	return result
}

