package main

func longestConsecutive(nums []int) int {
	ret := 0
	count := map[int]int{}
	parent := map[int]int{}
	for _, num := range nums {
		if _, ok := parent[num + 1]; ok {
			parent[num] = parent[num + 1]
		} else {
			parent[num] = num
		}

		if _, ok := parent[num - 1]; ok {
			parent[num - 1] = num
		}
	}

	for _, v := range parent {
		for v != parent[v] {
			v = parent[v]
		}
		count[v]++
		if count[v] > ret {
			ret = count[v]
		}
	}

	return ret
}
