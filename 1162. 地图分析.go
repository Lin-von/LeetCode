package main

import "math"

type V struct {
	x int
	y int
}

func maxDistance(grid [][]int) int {
	q := []V{}
	n := len(grid)
	if n == 0 {
		return -1
	}
	inq := make([][]bool, n)
	d := make([][]int, n)
	for i := 0; i < n ; i++ {
		inq[i] = make([]bool, n)
		d[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				d[i][j] = 0
				q = append(q, V{i,j})
				inq[i][j] = true
			} else {
				d[i][j] = math.MaxInt32
			}
		}
	}

	nx := []int{-1, 0, 1, 0}
	ny := []int{0, -1, 0, 1}

	for len(q) != 0 {
		v := q[0]
		q = q[1:len(q)]
		inq[v.x][v.y] = false
		for i := 0; i < 4; i++ {
			x := v.x + nx[i]
			y := v.y + ny[i]
			if x < 0 || y < 0 || x >= n || y >= n {
				continue
			}
			if d[x][y] > d[v.x][v.y] + 1 {
				d[x][y] = d[v.x][v.y] + 1
				if !inq[x][y] {
					nv := V{x, y}
					q = append(q, nv)
					inq[x][y] = true
				}
			}
		}
	}

	ret := -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && d[i][j] > ret {
				ret = d[i][j]
			}
		}
	}

	if ret == math.MaxInt32 {
		return -1
	} else {
		return ret
	}
}
