package main

func Min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}

	dp[0][0] = 0
	for i := 0; i <= m; i ++ {
		dp[i][0] = i
	}

	for i := 0; i <= n; i ++ {
		dp[0][i] = i
	}

	for i := 0; i < m; i ++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i + 1][j + 1] = dp[i][j]
			} else {
				dp[i + 1][j + 1] = 1 + Min(Min(dp[i + 1][j], dp[i][j + 1]), dp[i][j])
			}
		}
	}

	return dp[m][n]
}