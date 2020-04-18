package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func candy(ratings []int) int {
	ret := make([]int, len(ratings))
	ret[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i - 1] {
			ret[i] = ret[i - 1] + 1
		} else {
			ret[i] = 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i + 1] {
			ret[i] = max(ret[i], ret[i + 1] + 1)
		}
	}
	sum := 0
	for _, v := range ret {
		sum += v
	}
	return sum
}
