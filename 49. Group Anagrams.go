package main


func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int]int, 0)
	ret := [][]string{}
	for _, s := range strs {

		tmp := [26]int{}

		for i := 0; i < len(s); i++ {
			tmp[s[i] - 'a'] ++
		}

		if _, ok := m[tmp]; ok {
			ret[m[tmp]] = append(ret[m[tmp]], s)
		} else {
			tmpSlice := []string{}
			tmpSlice = append(tmpSlice, s)
			ret = append(ret, tmpSlice)
			m[tmp] = len(ret) - 1
		}
	}

	return ret
}
