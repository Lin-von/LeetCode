package main

func isNumber(s string) bool {
	word := map[byte]int{
		' ':0, '+':1, '-':1, '0':2, '1':2, '2':2, '3':2, '4':2, '5':2, '6':2, '7':2,
		'8':2, '9':2, '.':3, 'e':4,
	}

	state := [][]int{
		{ 0, 1, 6, 2,-1},
		{-1,-1, 6, 2,-1},
		{-1,-1, 3,-1,-1},
		{ 8,-1, 3,-1, 4},
		{-1, 7, 5,-1,-1},
		{ 8,-1, 5,-1,-1},
		{ 8,-1, 6, 3, 4},
		{-1,-1, 5,-1,-1},
		{ 8,-1,-1,-1,-1},
	}

	tmp := 0
	for i := 0; i < len(s); i++ {
		c, ok := word[s[i]]
		if !ok {
			return false
		}
		tmp = state[tmp][c];
		if tmp < 0 {
			return false
		}
	}
	return tmp == 3 || tmp == 5 || tmp == 6 || tmp == 8
}
