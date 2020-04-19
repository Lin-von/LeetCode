package main

func dfs(root *TreeNode, layer int, ret *[]int)  {
	if root == nil {
		return
	}

	if len(*ret) < layer {
		*ret = append(*ret, root.Val)
	} else {
		tmp := *ret
		tmp[layer - 1] = root.Val
	}

	dfs(root.Left, layer + 1, ret)
	dfs(root.Right, layer + 1, ret)
}

func rightSideView(root *TreeNode) []int {
	ret := []int{}
	dfs(root, 1, &ret)
	return ret
}
