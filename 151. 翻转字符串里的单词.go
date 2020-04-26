package main

func reverseWords(s string) string {
	ret := ""
	pos := len(s) - 1

	for pos >= 0 {
		for pos >= 0 && s[pos] == ' ' {
			pos --
		}
		if pos < 0 {
			break
		}
		i := pos
		for ; i >= 0 && s[i] != ' '; i--{}
		ret += s[i + 1:pos + 1]
		pos = i
		ret += " "
	}
	if ret == "" {
		return ret
	}
	return ret[:len(ret) - 1]
}
