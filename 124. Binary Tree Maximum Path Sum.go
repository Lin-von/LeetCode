package main

import "math"

func max(i ,j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

var ret int

func dfs(root *TreeNode) int {

	if root == nil {
		return 0
	}

	ret = max(ret, root.Val)

	l := max(dfs(root.Left), 0)
	r := max(dfs(root.Right), 0)

	ret = max(l + r + root.Val, ret)

	return max(l, r) + root.Val
}

func maxPathSum(root *TreeNode) int {
	ret = math.MinInt32
	dfs(root)
	return ret
}
