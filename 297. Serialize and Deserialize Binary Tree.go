package main

import (
	"strconv"
	"strings"
)

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

func ser(root *TreeNode, s *string)  {
	if root == nil {
		*s += "nil,"
	} else {
		*s += strconv.Itoa(root.Val) + ","
		ser(root.Left, s)
		ser(root.Right, s)
	}
}
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ret := ""
	ser(root, &ret)
	return ret
}

func deser(root *TreeNode, pos *int, arr []string, n int) {
	if *pos >= n {
		return
	}
	if arr[*pos] == "nil" {
		root.Left = nil
	} else {
		root.Left = new(TreeNode)
		root.Left.Val, _ = strconv.Atoi(arr[*pos])
		*pos ++
		deser(root.Left, pos, arr, n)
	}

	*pos ++
	if arr[*pos] == "nil" {
		root.Right = nil
	} else {
		root.Right = new(TreeNode)
		root.Right.Val, _ = strconv.Atoi(arr[*pos])
		*pos ++
		deser(root.Right, pos, arr, n)
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	if arr[0] == "nil" {
		return nil
	}
	ret := new(TreeNode)
	ret.Val, _ = strconv.Atoi(arr[0])
	pos := 1
	n := len(arr)
	deser(ret, &pos, arr, n)
	return ret
}

