package main

func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(p) > len(s) {
		return []int{}
	}
	cp := [26]int{}
	cs := [26]int{}
	ret := []int{}
	l := len(p)
	for i := 0; i < l; i ++ {
		cp[p[i] - 'a'] ++
		cs[s[i] - 'a'] ++
	}

	if cp == cs{
		ret = append(ret, 0)
	}

	pos := l
	for pos < len(s) {
		cs[s[pos - l] - 'a'] --
		cs[s[pos] - 'a'] ++
		pos ++
		if cp == cs{
			ret = append(ret, pos - l)
		}

	}
	return ret
}
