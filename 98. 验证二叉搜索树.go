package main

import "math"

func helper(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return helper(root.Left, root.Val, min) && helper(root.Right, max, root.Val)
}
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MaxInt64, math.MinInt64)
}