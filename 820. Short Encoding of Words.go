package main

type RTrie struct {
	isTail bool
	Next [26]*RTrie
}

func minimumLengthEncoding(words []string) int {
	r := new(RTrie)
	m := map[*RTrie]int{}
	for i := 0; i < len(words); i++ {
		tmp := r
		newmake := false
		for j := len(words[i]) - 1; j >= 0; j-- {
			pos := words[i][j] - 'a'
			if tmp.Next[pos] == nil {
				newmake = true
				tmp.Next[pos] = new(RTrie)
			}
			tmp.isTail = false
			tmp = tmp.Next[pos]
		}
		if newmake {
			tmp.isTail = true
		}
		m[tmp] = i
	}

	ret := 0
	for k, v := range m {
		if k.isTail {
			ret += len(words[v]) + 1
		}
	}
	return ret
}