package main

func hasGroupsSizeX(deck []int) bool {
	m := map[int]int{}
	for _, v := range deck {
		m[v] ++
	}

	tmp := m[deck[0]]
	for _, v := range m {
		if v < tmp {
			tmp = v
		}

		if v < 2 {
			return false
		}
	}

	for i := 2; i <= tmp; i++ {
		ok := true
		for _, v := range m {
			if v % i != 0 {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
