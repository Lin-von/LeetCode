package main

func replaceSpace(s string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ret += "%20"
		} else {
			ret += string(s[i])
		}
	}
	return ret
}
