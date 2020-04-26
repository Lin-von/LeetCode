package main

func numRookCaptures(board [][]byte) int {
	p, q := 0, 0
	for i := 0; i < len(board); i++ {
		found := false
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				p = i
				q = j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	ret := 0
	for i := p - 1; i >= 0; i-- {
		if board[i][q] == 'p' {
			ret ++
			break
		}
		if board[i][q] == 'B' {
			break
		}
	}

	for i := p + 1; i < len(board); i++ {
		if board[i][q] == 'p' {
			ret ++
			break
		}
		if board[i][q] == 'B' {
			break
		}
	}

	for i := q - 1; i >= 0; i-- {
		if board[p][i] == 'p' {
			ret ++
			break
		}
		if board[p][i] == 'B' {
			break
		}
	}

	for i := q + 1; i < len(board[0]); i++ {
		if board[p][i] == 'p' {
			ret ++
			break
		}
		if board[p][i] == 'B' {
			break
		}
	}
	return ret
}
