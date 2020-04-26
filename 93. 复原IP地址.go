package main

import (
	"strconv"
)

var res []string
var ip string
var ret []string
func isValid(s string) bool {
	n := len(s)
	if n > 3 {
		return false
	}
	if s[0] == '0' {
		return n == 1
	} else {
		a, _ := strconv.Atoi(s)
		return a <= 255
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func update(pos int)  {
	if pos < len(ip) && isValid(ip[pos:]) {
		tmp := ""
		for _, s := range res {
			tmp += s
			tmp += "."
		}
		tmp += ip[pos:]
		ret = append(ret, tmp)
	}
}

func dfs(pos, cnt int)  {
	n := len(ip)
	maxr := min(n, pos + 3)

	for i := pos + 1; i <= maxr; i ++ {
		tmp := ip[pos:i]
		if isValid(tmp) {
			res = append(res, tmp)
			if cnt == 2 {
				update(i)
			} else {
				dfs(i, cnt + 1)
			}
			res = res[:len(res) - 1]
		}
	}
}

func restoreIpAddresses(s string) []string {
	res = []string{}
	ret = []string{}
	ip = s
	dfs(0, 0)
	return ret
}
