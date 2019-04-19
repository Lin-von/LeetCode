package leetcode
func lengthOfLongestSubstring(s string) int {
	var r [128]int
	for i := 0; i < len(r); i++ {
		r[i] = -1
	}

	/*用pos维护无重复字符串的起点，找到重复字符就更新map，实时计算长度*/
	pos := -1
	result := 0
	for i := 0; i < len(s); i++ {
		pos = max(pos, r[s[i]])
		result = max(result, i - pos)
		r[s[i]] = i

	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	} else
	{
		return b
	}
}