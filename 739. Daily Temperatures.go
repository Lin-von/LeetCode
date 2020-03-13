package main

func dailyTemperatures(T []int) []int {
	stack := []int{}
	ret := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack) - 1]] {
			pos := stack[len(stack) - 1]
			ret[pos] = i - pos
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		pos := stack[len(stack) - 1]
		ret[pos] = 0
		stack = stack[:len(stack) - 1]
	}
	return ret
}
