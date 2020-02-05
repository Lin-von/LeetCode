package main

// 双递归解法，递归某个节点为起始节点的所有路径和递归所有节点为起始节点
//func Scan(root *TreeNode, sum int) int {
//	if root == nil {
//		return 0
//	}
//
//	cnt := 0
//	if root.Val == sum {
//		cnt++
//	}
//
//	return cnt + Scan(root.Left, sum - root.Val) + Scan(root.Right, sum - root.Val)
//}
//
//func pathSum(root *TreeNode, sum int) int {
//	if root == nil {
//		return 0
//	}
//
//	return Scan(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
//}

// 前缀和解法，树的路径拆成线段，用map维护某个路径和下有多少个区间符合条件，在其基础上递归
func Scan(root *TreeNode, preSum map[int]int, sum int, pathSum int) int {
	if root == nil {
		return 0
	}

	pathSum += root.Val
	cnt := preSum[pathSum - sum]
	preSum[pathSum]++
	cnt = cnt + Scan(root.Left, preSum, sum, pathSum) + Scan(root.Right, preSum, sum, pathSum)
	preSum[pathSum]--
	return cnt
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	preSum := map[int]int{0:1}

	return Scan(root, preSum, sum, 0)
}