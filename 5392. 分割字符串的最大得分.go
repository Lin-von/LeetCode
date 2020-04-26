package main

func maxScore(s string) int {
	ret := 0
	c1 := make([]int, len(s))
	c2 := make([]int, len(s))
	cnt0 := 0
	cnt1 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			cnt0 ++
		}
		c1[i] = cnt0
		if s[len(s) - 1 - i] == '1' {
			cnt1++
		}
		c2[len(s) - 1 - i] = cnt1
	}

	for i := 0; i < len(s) - 1; i++ {
		if c1[i] + c2[i + 1] > ret {
			ret = c1[i] + c2[i + 1]
		}
	}
	return ret
}
