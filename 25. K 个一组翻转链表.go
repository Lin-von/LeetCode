package main

func splitList(head *ListNode, k int) (*ListNode, bool) {
	if head == nil {
		return nil, false
	}

	p := head
	for k > 1 && p.Next != nil {
		p = p.Next
		k --
	}

	ret := p.Next
	p.Next = nil
	return ret, k == 1
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pre := head
	next := head.Next
	head.Next = nil
	for next != nil {
		tmp := next.Next
		next.Next = pre
		pre = next
		next = tmp
	}

	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := new(ListNode)
	pre.Next = head

	p := pre
	for {
		next, ok := splitList(head, k)
		if ok {
			pre.Next = reverseList(head)
			for pre.Next != nil {
				pre = pre.Next
			}
			head = next
		} else {
			pre.Next = head
			break
		}
	}
	return p.Next
}
