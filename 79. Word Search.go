package main

var w string
var b [][]byte

func dfs(i, j, m, n, index int, f [][]bool) bool {
	if i < 0 || i >= m || j < 0 || j >= n {
		return false
	}

	if f[i][j] {
		return false
	}

	if b[i][j] != w[index] {
		return false
	}

	if index == len(w) - 1 {
		return true
	}

	f[i][j] = true
	ret := dfs(i - 1, j, m, n, index + 1, f) || dfs(i + 1, j, m, n, index + 1, f) || dfs(i, j - 1, m, n, index + 1, f) || dfs(i, j + 1, m, n, index + 1, f)
	f[i][j] = false
	return ret

}

func exist(board [][]byte, word string) bool {
	w = word
	b = board
	m := len(board)
	n := len(board[0])
	f := make([][]bool, m)
	for i := 0; i < m; i++ {
		f[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && dfs(i, j, m, n, 0, f) {
				return true
			}
		}
	}
	return false
}
