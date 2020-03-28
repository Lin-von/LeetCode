package main

func isMirror(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	if l.Val == r.Val {
		return isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
	}

	return false
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}