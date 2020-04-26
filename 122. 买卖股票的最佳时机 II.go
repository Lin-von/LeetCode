package main

func maxProfit(prices []int) int {
	sum := 0
	pos := 0
	for pos != len(prices) {
		for pos + 1 < len(prices) && prices[pos] > prices[pos + 1] {
			pos ++
		}
		if pos + 1 == len(prices) {
			break
		}
		tmp := prices[pos]
		for pos + 1 < len(prices) && prices[pos] < prices[pos + 1] {
			pos ++
		}
		sum += prices[pos] - tmp
	}
	return sum
}
