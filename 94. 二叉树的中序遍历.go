package main

func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			ret = append(ret, cur.Val)
			cur = cur.Right
		} else {
			tmp := cur.Left
			for tmp.Right != nil {
				tmp = tmp.Right
			}
			tmp.Right = cur
			cur = cur.Left
			tmp.Right.Left = nil
		}
	}
	return ret
}