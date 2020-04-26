package main

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepth(root *TreeNode, num *int) int {
	if root == nil {
		return 0
	}
	ld := maxDepth(root.Left, num)
	rd := maxDepth(root.Right, num)
	*num = Max(*num, ld + rd)
	return Max(ld, rd) + 1
}


func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxNum := 0
	maxDepth(root, &maxNum)

	return maxNum

}
