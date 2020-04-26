package main

import "crypto/md5"

// 递归解法
//func Convert(root *TreeNode, sum *int) {
//	if root == nil {
//		return
//	}
//
//	Convert(root.Right, sum)
//	root.Val += *sum
//	*sum = root.Val
//	Convert(root.Left, sum)
//}
//
//func convertBST(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	sum := 0
//	Convert(root, &sum)
//	return root
//}

// Morris遍历，但为啥空间比递归还大？？
func Convert(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	p := root.Right
	if p == nil {
		root.Val += *sum
		*sum = root.Val
		Convert(root.Left, sum)
	} else {
		for p.Left != nil && p.Left != root {
			p = p.Left
		}
		if p.Left == nil {
			p.Left = root
			Convert(root.Right, sum)
		} else {
			p.Left = nil
			root.Val += *sum
			*sum = root.Val
			Convert(root.Left, sum)
		}
	}
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	sum := 0
	Convert(root, &sum)
	return root
}
