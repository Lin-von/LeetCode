package main

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func largestRectangleArea(heights []int) int {
	ret := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(heights); i++ {
		for stack[len(stack) - 1] != -1 && heights[stack[len(stack) - 1]] > heights[i] {
			tmp := heights[stack[len(stack) - 1]]
			if tmp * (i - stack[len(stack) - 2] - 1) > ret {
				ret = tmp * (i - stack[len(stack) - 2] - 1)
			}
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	for stack[len(stack) - 1] != -1 {
		tmp := heights[stack[len(stack) - 1]]
		if tmp * (len(heights) - stack[len(stack) - 2] - 1) > ret {
			ret = tmp * (len(heights) - stack[len(stack) - 2] - 1)
		}
		stack = stack[:len(stack) - 1]
	}

	return ret
}


func maximalRectangle(matrix [][]byte) int {
	ret := 0
	m := len(matrix)
	if (m == 0) {
		return ret
	}
	n := len(matrix[0])
	dp := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[j] ++
			} else {
				dp[j] = 0
			}
		}
		ret = max(ret, largestRectangleArea(dp))
	}

	return ret
}
