package main


// DP dp[i] = Max(dp[i-2] + nums[i], num[i - 1])
// 使用两个指针分别表示dp[i-2] dp[i-1]即可
func rob(nums []int) int {
	fst := 0
	snd := 0
	sum := 0
	for _, num := range nums {
		if fst + num > snd {
			sum = fst + num
		} else {
			sum = snd
		}

		fst = snd
		snd = sum
	}

	return sum
}
