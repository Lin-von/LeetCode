package main

var cnt int
var f [][]int

func dfs(i, j, k, m ,n int) {
	if i < 0 || i >= m || j < 0 || j >= n || f[i][j] > k {
		return
	}

	cnt ++
	f[i][j] = k + 1

	dfs(i + 1, j, k, m, n)
	dfs(i - 1, j, k, m, n)
	dfs(i, j - 1, k, m, n)
	dfs(i, j + 1, k, m, n)
}

func movingCount(m int, n int, k int) int {
	cnt = 0
	f = make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}


	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p := i
			q := j
			for p != 0 {
				f[i][j] += p % 10
				p /= 10
			}
			for q != 0 {
				f[i][j] += q % 10
				q /= 10
			}
		}
	}

	dfs(0, 0, k, m, n)
	return cnt
}
