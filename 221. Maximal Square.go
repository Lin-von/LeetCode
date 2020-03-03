package main

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if (m == 0) {
		return 0
	}
	n := len(matrix[0])
	dp := make([]int, n + 1)
	last := 0
	ret := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if matrix[i - 1][j - 1] == '1' {
				dp[j] = Min(dp[j], Min(dp[j - 1], last)) + 1
				ret = Max(ret, dp[j])
			} else {
				dp[j] = 0
			}
			last = tmp
		}
	}
	return ret * ret
}
