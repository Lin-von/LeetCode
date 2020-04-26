package main

func scanByLevel(arr *[][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(*arr) <= level {
		*arr = append(*arr, []int{})
	}

	tmp := *arr
	tmp[level] = append(tmp[level], root.Val)
	scanByLevel(arr, root.Left, level + 1)
	scanByLevel(arr, root.Right, level + 1)
}


func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	scanByLevel(&ret, root, 0)
	for i := 0; i < len(ret); i++ {
		if i & 1 == 1 {
			tmp := ret[i]
			l := 0
			r := len(tmp) - 1
			for l < r {
				tmp[l], tmp[r] = tmp[r], tmp[l]
				l ++
				r --
			}
		}
	}
	return ret
}
