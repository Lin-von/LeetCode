package main

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	var dp [][]bool
	dp = make([][]bool, len(s) + 1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p) + 1)
	}

	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		if p[i - 1] == '*' {
			dp[0][i] = dp[0][i - 2]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j<= len(p); j++ {
			if s[i - 1] == p[j - 1] || p[j - 1] == '.' {
				dp[i][j] = dp[i - 1][j - 1]
			}

			// 状态转移主要处理*符号
			if p[j - 1] == '*' {
				// *号前一个与文本匹配，则有三种可能路径：空匹配、多字符匹配、单字符匹配
				if s[i - 1] == p[j - 2] || p[j - 2] == '.' {
					dp[i][j] = dp[i][j] || dp[i][j - 2] || dp[i - 1][j] || dp[i][j - 1]
				}
				// 未匹配，用*向前移两个字符的结果
				if s[i - 1] != p[j - 2] {
					dp[i][j] = dp[i][j] || dp[i][j - 2]
				}
			}

		}

	}

	return dp[len(s)][len(p)]
}
