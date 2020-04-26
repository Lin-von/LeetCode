package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s) + 1)
	dp[0] = true

	dict := map[string]bool{}
	for _, s := range wordDict {
		dict[s] = true
	}

	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}
