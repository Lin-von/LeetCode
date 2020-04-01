package main

func pathScan(root *TreeNode, sum int, cur []int,  arr *[][]int)  {
	if root == nil {
		return
	}

	cur = append(cur, root.Val)

	if root.Left == nil && root.Right == nil && root.Val == sum {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*arr = append(*arr, tmp)
		return
	}

	pathScan(root.Left, sum - root.Val, cur, arr)
	pathScan(root.Right, sum - root.Val, cur, arr)
}
func pathSum(root *TreeNode, sum int) [][]int {
	ret := [][]int{}
	pathScan(root, sum, []int{}, &ret)
	return ret
}