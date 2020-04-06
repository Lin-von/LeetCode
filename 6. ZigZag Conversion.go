package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	cur := make([]string, numRows)
	pos := 0
	step := -1
	for i := 0; i < len(s); i++ {
		cur[pos] += string(s[i])
		if pos == numRows - 1 || pos == 0 {
			step *= -1
		}
		pos += step
	}
	ret := ""
	for i := 0; i < numRows; i++ {
		ret += cur[i]
	}
	return ret
}
