package main

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := root.Left
	r := root.Right

	if l != nil {
		for l.Right != nil {
			l = l.Right
		}
		if l.Val >= root.Val {
			return false
		}
	}

	if r != nil {
		for r.Left != nil {
			r = r.Left
		}
		if r.Val <= root.Val {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}