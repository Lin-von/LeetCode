package main

func gameOfLife(board [][]int)  {
	n := len(board)
	if n == 0 {
		return
	}
	m := len(board[0])
	step := []int{-1, 0, 1}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cnt := 0
			for nx := 0; nx < 3; nx ++ {
				for ny := 0; ny < 3; ny ++ {
					if nx == 1 && ny == 1 {
						continue
					}
					x := i + step[nx]
					y := j + step[ny]
					if x < 0 || y < 0 || x >= n || y >= m {
						continue
					}
					if board[x][y] == 1 || board[x][y] == -1 {
						cnt ++
					}
				}
			}
			if board[i][j] == 1 {
				if cnt < 2 || cnt > 3 {
					board[i][j] = -1
				}
			} else {
				if cnt == 3 {
					board[i][j] = 2
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}
