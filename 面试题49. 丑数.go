package main

func min(a, b, c int) int {
	tmp := a
	if b < a {
		tmp = b
	}
	if tmp < c {
		return tmp
	} else {
		return c
	}
}
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	pos2, pos3, pos5 := 0, 0, 0
	for i := 1; i < n; i++ {
		a := dp[pos2] * 2
		b := dp[pos3] * 3
		c := dp[pos5] * 5
		dp[i] = min(a, b, c)
		if dp[i] == a {
			pos2 ++
		}
		if dp[i] == b {
			pos3 ++
		}
		if dp[i] == c {
			pos5 ++
		}
	}
	return dp[n - 1]
}
