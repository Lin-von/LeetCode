package main

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	ret := []int{}
	if n == 0 {
		return ret
	}
	m := len(matrix[0])

	cnt := 0
	l := 0
	r := m - 1
	u := 0
	d := n - 1
	for cnt < n * m {
		for i := l; i <= r; i++ {
			ret = append(ret, matrix[u][i])
			cnt ++
			if cnt == n * m {
				return ret
			}
		}
		u ++
		for i := u; i <= d; i++ {
			ret = append(ret, matrix[i][r])
			cnt ++
			if cnt == n * m {
				return ret
			}
		}
		r --
		for i := r; i >= l; i-- {
			ret = append(ret, matrix[d][i])
			cnt ++
			if cnt == n * m {
				return ret
			}
		}
		d --
		for i := d; i >= u; i-- {
			ret = append(ret, matrix[i][l])
			cnt ++
			if cnt == n * m {
				return ret
			}
		}
		l ++
	}

	return ret
}
