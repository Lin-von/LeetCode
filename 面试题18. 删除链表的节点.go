package main

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	pre := head
	cur := head.Next
	for cur.Val != val {
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = cur.Next
	return head
}
