package main

func findDiagonalOrder(nums [][]int) []int {
	ret := []int{}

	tmp := [][]int{}

	for i := 0; i < len(nums); i ++ {
		for j := 0; j < len(nums[i]); j++ {
			if len(tmp) == i + j {
				tmp = append(tmp, []int{})
			}
			tmp[i + j] = append([]int{nums[i][j]}, tmp[i + j]...)
		}
	}

	for i := 0; i < len(tmp); i++ {
		ret = append(ret, tmp[i]...)
	}

	return ret

}
