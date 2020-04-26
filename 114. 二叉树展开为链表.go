package main

func flatten(root *TreeNode)  {
	cur := root
	for cur != nil {
		tmp := cur.Left
		if tmp == nil {
			cur = cur.Right
		} else {
			pos := tmp
			for tmp.Right != nil {
				tmp = tmp.Right
			}
			tmp.Right = cur.Right
			cur.Right = pos
			cur.Left = nil
			cur = cur.Right
		}
	}
}
