package main

import "strconv"

func getPermutation(n int, k int) string {
	k --
	tmp := 1
	pos := 1
	ret := ""

	for pos < n {
		tmp *= pos
		pos ++
	}

	pos --
	used := make([]bool, n)

	for pos > 0 {
		cur := k / tmp
		for i := 0; i < n; i++ {
			if !used[i] {
				if cur == 0 {
					used[i] = true
					ret += strconv.Itoa(i + 1)
					break
				} else {
					cur --
				}
			}
		}
		k %= tmp
		tmp /= pos
		pos --
	}

	for i := 0; i < n; i++ {
		if !used[i] {
			ret += strconv.Itoa(i + 1)
		}
	}

	return ret

}