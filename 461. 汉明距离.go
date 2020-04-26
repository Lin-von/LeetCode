package main

func hammingDistance(x int, y int) int {
	x = x ^ y
	y = 0
	for x != 0 {
		if x & 1 == 1 {
			y++
		}
		x = x >> 1
	}
	return y
}
