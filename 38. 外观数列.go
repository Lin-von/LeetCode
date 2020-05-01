package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)
	ret := ""
	var tmp uint8
	cnt := 0
	pos := 0
	for pos < len(s) {
		tmp = s[pos]
		cnt = 1
		for pos + 1 < len(s) && s[pos] == s[pos + 1] {
			cnt ++
			pos ++
		}
		ret += strconv.Itoa(cnt)
		ret += string(tmp)
		pos ++
	}
	return ret
}
