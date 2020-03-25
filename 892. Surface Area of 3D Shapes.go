package main

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func surfaceArea(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	ret := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				continue
			}
			ret += 2
			ret += grid[i][j] * 4
			if j > 0 {
				ret -= min(grid[i][j - 1], grid[i][j]) * 2
			}

			if i > 0 {
				ret -= min(grid[i - 1][j], grid[i][j]) * 2
			}

		}
	}
	return ret
}
