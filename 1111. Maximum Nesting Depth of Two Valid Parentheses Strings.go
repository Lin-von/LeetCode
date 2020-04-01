package main

func maxDepthAfterSplit(seq string) []int {
	ret := []int{}
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			ret = append(ret, i & 1)
		} else {
			ret = append(ret, i & 1 ^ 1)
		}
	}
	return ret
}
