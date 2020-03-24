package main

var pos int
var pre []int
var idx map[int]int

func build(l, r int) *TreeNode {
	if l == r {
		return nil
	}

	ret := new(TreeNode)
	ret.Val = pre[pos]
	pos ++

	ret.Left = build(l, idx[ret.Val])
	ret.Right = build(idx[ret.Val] + 1, r)
	return ret
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	pos = 0
	pre = preorder
	idx = map[int]int{}
	for i := 0; i < len(inorder); i++ {
		idx[inorder[i]] = i
	}

	return build(0, len(pre))
}
