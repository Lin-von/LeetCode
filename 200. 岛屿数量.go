package main

var f [][]bool
var g [][]byte
var n, m int

func dfs(i, j int)  {
	if i < 0 || i >= n || j < 0 || j >= m {
		return
	}

	if f[i][j] || g[i][j] != '1' {
		return
	}

	f[i][j] = true
	dfs(i - 1, j)
	dfs(i + 1, j)
	dfs(i, j - 1)
	dfs(i, j + 1)
}

func numIslands(grid [][]byte) int {
	ret := 0
	n = len(grid)
	if n == 0 {
		return 0
	}
	
	m = len(grid[0])
	g = grid
	f = make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !f[i][j] && grid[i][j] == '1' {
				dfs(i, j)
				ret++
			}
		}
	}
	return ret
}
