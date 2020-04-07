package main

func scan(m [][]int, s, n int) {
	for i := s; i <= n - 1 - s - 1; i++ {
		m[s][i], m[n - i - 1][s] = m[n - i - 1][s], m[s][i]
		m[n - i - 1][s], m[n - 1 - s][n - 1 - i] = m[n - 1 - s][n - 1 - i], m[n - i - 1][s]
		m[n - 1 - s][n - 1 - i], m[i][n - 1 - s] = m[i][n - 1 - s], m[n - 1 - s][n - 1 - i]
	}
}

func rotate(matrix [][]int)  {
	n := len(matrix)
	if n == 0 {
		return
	}
	for i := 0; i <= n / 2; i++ {
		scan(matrix, i, n)
	}
}