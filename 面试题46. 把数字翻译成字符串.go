package main

func dfs(cur string, num int, f map[string]bool)  {
	if num == 0 {
		f[cur] = true
		return
	}

	tmpNum := num % 10
	tmpStr := string(rune(tmpNum + 97))
	dfs(tmpStr + cur, num / 10, f)

	if num / 10 > 0 && num % 100 < 26 && num % 100 > 9 {
		tmpNum = num % 100
		tmpStr = string(rune(tmpNum + 97))
		dfs(tmpStr + cur, num / 100, f)
	}
}

func translateNum(num int) int {
	f := map[string]bool{}
	dfs("", num, f)
	return len(f)
}
