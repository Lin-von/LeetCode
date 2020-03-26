package main

func printNumbers(n int) []int {
	ret := []int{}
	mmax := 1
	tmp := 10
	for n != 0 {
		if (n & 1 == 1) {
			mmax *= tmp
		}
		tmp *= tmp
		n >>= 1
	}
	for i := 1; i < mmax; i++ {
		ret = append(ret, i)
	}
	return ret
}
