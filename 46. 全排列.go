package main

var arr []int

func scan(cur []int, ret *[][]int)  {
	if len(cur) == len(arr) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}

	for i := 0; i < len(arr); i++ {
		have := false
		for _, num := range cur {
			if num == arr[i] {
				have = true
				break
			}
		}
		if have {
			continue
		}

		tmp := append(cur, arr[i])
		scan(tmp, ret)
	}

}

func permute(nums []int) [][]int {
	arr = nums
	c := []int{}
	ret := [][]int{}
	scan(c, &ret)
	return ret
}