package main

import "math"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func updateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				ret[i][j] = 0
			} else {
				ret[i][j] = math.MaxInt32
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i - 1 >= 0 {
				ret[i][j] = min(ret[i - 1][j] + 1, ret[i][j])
			}
			if j - 1 >= 0 {
				ret[i][j] = min(ret[i][j - 1] + 1, ret[i][j])
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j :=  m - 1; j >= 0; j-- {
			if i + 1 < n {
				ret[i][j] = min(ret[i + 1][j] + 1, ret[i][j])
			}
			if j + 1 < m {
				ret[i][j] = min(ret[i][j + 1] + 1, ret[i][j])
			}
		}
	}

	return ret
}
