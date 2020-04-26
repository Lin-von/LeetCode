package main

var pos int
var pre []int
var idx map[int]int

func build(l, r int) *TreeNode {
	if l == r {
		return nil
	}

	tmp := TreeNode{}
	tmp.Val = pre[pos]

	pos ++
	tmp.Left = build(l, idx[tmp.Val])
	tmp.Right = build(idx[tmp.Val] + 1, r)
	return &tmp
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	pre = preorder
	pos = 0
	idx = map[int]int{}
	for i := 0; i < len(inorder); i++ {
		idx[inorder[i]] = i
	}

	return build(0, len(inorder))
}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(inorder) == 0 {
//		return nil
//	}
//
//	root := TreeNode{
//		Val:   preorder[0],
//		Left:  nil,
//		Right: nil,
//	}
//
//	for i := 0; i < len(inorder); i++ {
//		if root.Val == inorder[i] {
//			root.Left = buildTree(preorder[1:], inorder[0:i])
//			root.Right = buildTree(preorder[i + 1:], inorder[i + 1:])
//		}
//	}
//
//	return &root
//}