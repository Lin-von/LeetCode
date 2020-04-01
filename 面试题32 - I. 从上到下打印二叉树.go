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

func levelOrder(root *TreeNode) []int {
	rec := [][]int{}
	scanByLevel(&rec, root, 0)
	ret := []int{}
	for _, arr := range rec {
		ret = append(ret, arr...)
	}
	return ret
}
