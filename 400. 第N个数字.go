package main

import "strconv"

func findNthDigit(n int) int {
	num := 1
	cnt := 1
	n --
	for {
		if 9 * num * cnt > n {
			break
		}
		n -= 9 * num * cnt
		num ++
		cnt *= 10
	}
	tmp := cnt + n / num
	pos := n % num
	return int(strconv.Itoa(tmp)[pos] - '0')
}
