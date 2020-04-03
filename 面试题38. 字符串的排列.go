package main

import "sort"

func dfs(cur string, s string, f []bool, tmp map[string]bool)  {
	if len(cur) == len(s) {
		tmp[cur] = true
		return
	}

	for i := 0; i < len(s); i++ {
		if f[i] {
			continue
		}

		if i > 0 && s[i] == s[i - 1] && !f[i - 1] {
			continue
		}

		f[i] = true
		dfs(cur + string(s[i]), s, f, tmp)
		f[i] = false
	}
}

func permutation(s string) []string {
	f := make([]bool, len(s))
	ret := []string{}
	tmp := map[string]bool{}
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	s = string(t)
	dfs("", s, f, tmp)
	for k, _ := range tmp {
		ret = append(ret, k)
	}
	return ret
}
