package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//可以理解为分成两条路径比较，两条路径取左右子树时相反，比较好理解的方式就是复制一棵树镜像路径比较，实际上直接使用左右子树也相当于复制树来比较
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