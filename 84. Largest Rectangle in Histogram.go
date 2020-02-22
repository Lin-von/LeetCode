package main

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
