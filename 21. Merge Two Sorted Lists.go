package main

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var node ListNode
	p := &node

	for ; l1 != nil && l2 != nil; {
		if l1.Val > l2.Val {
			p.Next = l2
			l2 = l2.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}

	return node.Next
}
