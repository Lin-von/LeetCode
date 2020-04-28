package main

func isValidSudoku(board [][]byte) bool {
	clo := make([][]bool, 9)
	raw := make([][]bool, 9)
	rec := make([][]bool, 9)
	for i := 0; i < 9; i ++ {
		clo[i] = make([]bool, 9)
		raw[i] = make([]bool, 9)
		rec[i] = make([]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			tmp := board[i][j]
			if tmp == '.' {
				continue
			}
			num := tmp - '1'
			if raw[i][num] {
				return false
			} else {
				raw[i][num] = true
			}

			if clo[j][num] {
				return false
			} else {
				clo[j][num] = true
			}

			pos := i / 3 * 3 + j / 3
			if rec[pos][num] {
				return false
			} else {
				rec[pos][num] = true
			}
		}
	}
	return true
}
