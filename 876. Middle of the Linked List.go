package main

func middleNode(head *ListNode) *ListNode {
	p := head
	q := head.Next
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	if q != nil {
		p = p.Next
	}
	return p
}
