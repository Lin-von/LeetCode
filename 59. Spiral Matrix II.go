package main

func scan(m [][]int, s, n int, num *int) {
	for i := s; i <= n - s - 1; i++ {
		m[s][i] = *num
		*num ++
	}

	for i := s + 1; i <= n - s - 1; i ++  {
		m[i][n - s - 1] = *num
		*num ++
	}

	for i := n - s - 2; i >= s; i-- {
		m[n - s - 1][i] = *num
		*num ++
	}

	for i := n - s - 2; i > s; i-- {
		m[i][s] = *num
		*num ++
	}
}

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	num := 1
	for i := 0; i <= n / 2; i++ {
		scan(ret, i, n, &num)
	}
	return ret
}
