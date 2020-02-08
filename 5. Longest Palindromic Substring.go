package main


func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	dp := make([]int, n)
	ret := s[0:1]


	// 优化DP空间，二维DP表示的是i起始j结束是否为回文，实际dp[i][j]只用到dp[i+1][j-1]。
	// 因此优化为一维，DP[j]表示第i+1行状态，即一维DP随着运行时行数变化在不断更新，因此为了防止未使用的数据被覆盖，采用倒序遍历j
	for i := n - 1; i >= 0; i-- {
		for j:= n - 1; j >= i; j-- {
			if s[i] == s[j] && ( (i + 1 >= j - 1) || dp[j - 1] == 1 ) {
				dp[j] = 1
				if j - i + 1 > len(ret) {
					ret = s[i:j + 1]
				}
			} else {
				dp[j] = 0
			}
		}
	}

	return ret
}