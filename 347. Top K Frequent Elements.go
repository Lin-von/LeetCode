package main

func topKFrequent(nums []int, k int) []int {
	ret := []int{}
	// 第一步基础桶排序
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num] ++
	}

	// 将桶排序的结果再做桶排序，得到出现频率的数组
	reflc := map[int][]int{}
	for k, v := range cnt {
		reflc[v] = append(reflc[v], k)
	}

	// 基于题目给定K必定有效，从理论最大值N向下找存在的桶，将结果追加进去即可
	pos := len(nums)
	for k != 0 {
		if tmp, ok := reflc[pos]; ok {
			k -= len(tmp)
			ret = append(ret, tmp...)
		}
		pos --
	}
	return ret
}
