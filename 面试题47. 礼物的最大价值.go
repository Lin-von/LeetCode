package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxValue(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[0][i] = grid[0][i] + dp[0][i - 1]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = grid[i][0] + dp[i - 1][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
		}
	}

	return dp[n - 1][m - 1]

}
