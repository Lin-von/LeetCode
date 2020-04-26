package main

func swap(m [][]int, i, j, n int)  {
	m[i][j], m[j][n - 1 -i], m[n - 1 - i][n - 1 - j], m[n - 1 - j][i] = m[n - 1 - j][i], m[i][j], m[j][n - 1 -i], m[n - 1 - i][n - 1 - j]
}

func rotate(matrix [][]int)  {
	n := len(matrix[0])
	for i := 0; i < n / 2; i++ {
		for j := i; j < n - i - 1; j++ {
			swap(matrix, i, j, n)
		}
	}
}
