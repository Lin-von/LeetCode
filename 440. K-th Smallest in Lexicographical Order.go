package main

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func getCount(prefix, n int) int {
	a := prefix
	b := a + 1
	cnt := 0
	for a <= n {
		cnt += min(n + 1, b) - a
		a *= 10
		b *= 10
	}
	return cnt
}

func findKthNumber(n int, k int) int {
	prefix := 1
	cnt := 1
	for cnt < k {
		tmp := getCount(prefix, n)
		if cnt + tmp <= k {
			prefix ++
			cnt += tmp
		} else {
			prefix *= 10
			cnt ++
		}
	}
	return prefix
}
