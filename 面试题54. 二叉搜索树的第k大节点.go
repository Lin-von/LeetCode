package main

func dfs(root *TreeNode, ret *int, cur *int, tar int) {
	if root == nil || *cur == tar {
		return
	}
	dfs(root.Right, ret, cur, tar)
	*cur ++
	if *cur == tar {
		*ret = root.Val
		return
	}
	dfs(root.Left, ret, cur, tar)
}

func kthLargest(root *TreeNode, k int) int {
	cur := 0
	ret := 0
	dfs(root, &ret, &cur, k)
	return ret
}
