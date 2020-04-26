package main

var ret []string

func isValid(s string) bool {
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return true
}

func dfs(cur string, pos int, wl, wr int)  {
	if wl == 0 && wr == 0 && isValid(cur){
		ret = append(ret, cur)
	}

	for i := pos; i < len(cur); i++ {
		if i != pos && cur[i] == cur[i - 1] {
			continue
		}

		if cur[i] == '(' && wl > 0 {
			dfs(cur[0:i] + cur[i + 1:], i, wl - 1, wr)
		} else if cur[i] == ')' && wr > 0 {
			dfs(cur[0:i] + cur[i + 1:], i, wl, wr - 1)
		}
	}

}

func removeInvalidParentheses(s string) []string {
	wl := 0
	wr := 0
	for i := 0; i < len(s); i ++ {
		if s[i] == '(' {
			wl ++
		}
		if s[i] == ')' {
			if wl > 0 {
				wl --
			} else {
				wr ++
			}
		}
	}
	ret = []string{}
	dfs(s, 0, wl, wr)
	return ret
}
