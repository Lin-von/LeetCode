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

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	scanByLevel(&ret, root, 0)
	for i, a := range ret{
		if i & 1 == 1 {
			for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
				a[left], a[right] = a[right], a[left]
			}
		}
	}
	return ret
}