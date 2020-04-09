package main

func constructArr(a []int) []int {
	ret := make([]int, len(a))
	for i := 0; i < len(ret); i++ {
		ret[i] = 1
	}

	for i := 1; i < len(ret); i++ {
		ret[i] = ret[i - 1] * a[i - 1]
	}

	tmp := 1
	for i := len(ret) - 2; i >= 0; i -- {
		tmp *= a[i + 1]
		ret[i] *= tmp
	}
	return ret
}
