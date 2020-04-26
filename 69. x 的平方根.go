package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	var tar int64
	tar = int64(x)

	l := 2
	r := x / 2

	for l <= r {
		var tmp int64
		k := (l + r) >> 1
		tmp = int64(k * k)
		if tmp == tar {
			return k
		} else if tmp < tar {
			l = k + 1
		} else {
			r = k - 1
		}
	}

	return r
}
