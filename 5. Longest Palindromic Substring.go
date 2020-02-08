package main


func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	dp := [1000][1000]int{}
	ret := s[0:1]
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n ; j++ {

			if s[i] == s[j] && ( (i + 1 > j - 1) || dp[i + 1][j - 1] == 1 ) {
				dp[i][j] = 1
				if j - i + 1 > len(ret) {
					ret = s[i:j + 1]
				}
			} else {
				dp[i][j] = 0
			}

		}
	}
	return ret
}