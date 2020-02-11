package main

func longestValidParentheses(s string) int {
	l, r, max := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l ++
		} else {
			r ++
		}

		if l == r && l + r > max {
			max = l + r
		} else if r > l {
			l, r = 0, 0
		}
	}

	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			l ++
		} else {
			r ++
		}

		if l == r && l + r > max {
			max = l + r
		} else if r < l {
			l, r = 0, 0
		}
	}

	return max
}

