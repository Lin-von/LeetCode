package main

func isValid(s string) bool {
	if len(s) & 1 != 0 {
		return false
	}

	if s == "" {
		return true
	}

	stack := make([]uint8, len(s))
	pos := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack[pos] = s[i]
			pos ++
		} else {
			pos --
			if pos < 0 {
				return false
			}
			if s[i] == ')' && stack[pos] != '(' {
				return false
			}
			if s[i] == ']' && stack[pos] != '[' {
				return false
			}
			if s[i] == '}' && stack[pos] != '{' {
				return false
			}
		}
	}
	return pos == 0
}
