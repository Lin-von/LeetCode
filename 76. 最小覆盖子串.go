package main

// 滑动窗口，用一个哈希表来表示子串的字符及个数。
// 移动窗口右边界，统计字母个数，如果还未统计到子串的字母个数（即哈希表还大于0）就增加匹配长度。
// 注意本题匹配要求字母的个数也完全一致，因此在这种统计下匹配长度等于子串时就说明刚好满足条件
// 在窗口满足条件的情况下，更新结果，移动左边界，同理：将之前做统计减掉的个数加回来，如果统计的个数不为0了（统计到子串字母了），就减小匹配长度
func minWindow(s string, t string) string {
	l := 0
	p := 0
	q := 0
	h := [128]int{}
	ret := ""

	for i := 0; i < len(t); i++ {
		h[t[i]] += 1
	}

	for q < len(s) {
		h[s[q]] --
		if h[s[q]] >=0 {
			l ++
		}
		q++
		for l == len(t) {
			if ret == "" || q - p < len(ret) {
				ret = s[p:q]
			}
			h[s[p]] ++
			if h[s[p]] > 0 {
				l --
			}
			p ++
		}
	}
	return ret

}