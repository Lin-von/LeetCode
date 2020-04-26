package main

func longestPalindrome(s string) int {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]] ++
	}
	sum := 0
	tmp := false
	for _, v := range m {
		if v & 1 == 0 {
			sum += v
		} else {
			tmp = true
			sum += v - 1
		}
	}
	if tmp {
		return sum + 1
	} else {
		return sum
	}
}
