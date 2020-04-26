package main


func swapPairs(head *ListNode) *ListNode {
	p := head
	if p == nil {
		return nil
	}
	q := head.Next
	if q == nil {
		return p
	}
	next := q.Next
	q.Next = p
	p.Next = swapPairs(next)
	return q
}