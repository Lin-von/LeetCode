package main

var ret int
var f []bool
func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
func dfs(x ,y int, n, m, cur int)  {
	if cur > n {
		return
	}
	if cur >= m {
		ret ++
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == x && j == y {
				continue
			}
			if f[i * 3 + j] {
				continue
			}
			if abs(x - i) % 2 == 0 && abs(y - j) % 2 == 0 && !f[(x + i) / 2 * 3 + (j + y) / 2] {
				continue
			}
			f[i * 3 + j] = true
			dfs(i, j, n, m, cur + 1)
			f[i * 3 + j] = false
		}
	}
}

func numberOfPatterns(m int, n int) int {
	f = make([]bool, 9)
	ret = 0
	f[0] = true
	dfs(0, 0, n, m, 1)
	f[0] = false
	f[1] = true
	dfs(0, 1, n, m, 1)
	f[1] = false
	ret *= 4
	f[4] = true
	dfs(1, 1, n, m, 1)
	f[4] = false
	return ret
}
