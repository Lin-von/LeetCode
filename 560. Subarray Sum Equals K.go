package main

func subarraySum(nums []int, k int) int {
	pre := map[int]int{}
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	ret := 0
	pre[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ret += pre[sum - k]
		pre[sum] ++
	}
	return ret
}
