package main

func isSame(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil {
		return false
	}

	return A.Val == B.Val && isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return isSame(A, B) ||
		isSubStructure(A.Left, B) || isSubStructure(A.Right, B)

}
