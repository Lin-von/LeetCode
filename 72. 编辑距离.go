package main

func Min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

// DP[i][j]表示字符串1的i和2的j之间的编辑距离
// 相同时无需操作，不相同时的操作删、加、改可以分别转换为三个前置状态
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