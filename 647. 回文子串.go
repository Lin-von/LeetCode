package main

func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	sum := 1
	for i := 1; i < n; i++ {
		dp[i][i] = true
		sum ++
		if s[i] == s[i - 1] {
			sum ++
			dp[i - 1][i] = true
		}
		for j := 0; j < i - 1; j++ {
			if s[i] == s[j] && dp[j + 1][i - 1] {
				sum ++
				dp[j][i] = true
			}
		}
	}
	return sum
}

func Min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func countSubstrings(s string) int {
	t := "^#"
	for i := 0; i < len(s); i++ {
		t += string(s[i])
		t += "#"
	}
	t += "$"
	n := len(t)
	p := make([]int, n)
	maxRight := 0
	center := 0
	for i := 1; i < n - 1; i++ {
		if i < maxRight {
			p[i] = Min(maxRight - i, p[center * 2 - i])
		}
		for t[i - p[i] - 1] == t[i + p[i] + 1] {
			p[i] ++
		}
		if i + p[i] > maxRight {
			maxRight = i + p[i]
			center = i
		}
	}

	ret := 0
	for _, v := range p {
		ret += (v + 1) / 2
	}

	return ret
}

