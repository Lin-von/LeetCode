package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxScore(cardPoints []int, k int) int {
	if k == 0 {
		return 0
	}

	if len(cardPoints) == 1 {
		return cardPoints[0]
	}

	if k == 1 {
		return max(cardPoints[0], cardPoints[len(cardPoints) - 1])
	}

	if k >= len(cardPoints) {
		ret := 0
		for _, v := range cardPoints {
			ret += v
		}
		return ret
	}

	ret := 0
	tmp := 0
	for i := 0; i < k; i++ {
		ret	+= cardPoints[i]
		tmp = ret
	}

	for i := k - 1; i >= 0; i-- {
		tmp -= cardPoints[i]
		tmp += cardPoints[len(cardPoints) - k + i]
		ret = max(ret, tmp)
	}
	return ret
}
