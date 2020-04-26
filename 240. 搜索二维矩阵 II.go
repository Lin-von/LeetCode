package main

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	i := n - 1
	j := 0
	for j < m && i >= 0 {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			i --
		} else {
			j ++
		}
	}
	return false
}
