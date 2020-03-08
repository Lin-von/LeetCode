package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 对于每个节点都有两种状态： 选与不选。如果选了就不能选子节点；同样，对于其每个子节点也有两种状态，当不选当前节点时，结果就是左右各节点状态较大的那个之和
// 在搜索时一个顾虑点时，当前节点如果选了会对其父节点造成影响，因此我们从上向下搜索，对于dfs返回每个节点的两种状态，将当前节点对父节点的影响交由父节点来决定
func dfs(root *TreeNode) (child, father int) {
	if root == nil {
		return 0, 0
	}

	lc, lf := dfs(root.Left)
	rc, rf := dfs(root.Right)

	return max(lc, lf) + max(rc, rf), root.Val + lc + rc
}

func rob(root *TreeNode) int {
	return max(dfs(root))
}
