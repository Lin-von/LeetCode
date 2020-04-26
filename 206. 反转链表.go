package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head.Next
	c := head
	head.Next = nil
	for p != nil {
		n := p.Next
		p.Next = c
		c = p
		p = n
	}
	return c
}