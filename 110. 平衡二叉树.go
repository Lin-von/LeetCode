package main

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left)
	r := dfs(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if abs(l - r) > 1 {
		return -1
	} else {
		return max(l, r) + 1
	}
}

func isBalanced(root *TreeNode) bool {
	return dfs(root) != -1
}
