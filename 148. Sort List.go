package main

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

func splitList(l *ListNode, length int) *ListNode {
	if l == nil {
		return nil
	}

	for length > 1 && l.Next != nil {
		length --
		l = l.Next
	}

	ret := l.Next
	l.Next = nil
	return ret
}


func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 1
	f := head.Next
	for f != nil && f.Next != nil {
		length ++
		f = f.Next.Next
	}
	length <<= 1
	if f == nil {
		length --
	}

	ret := ListNode{}
	ret.Next = head

	for step := 1; step < length; step <<= 1 {
		prev := &ret
		next := ret.Next
		for next != nil {
			l1 := next
			l2 := splitList(l1, step)
			next = splitList(l2, step)
			prev.Next = mergeTwoLists(l1, l2)

			for prev.Next != nil {
				prev = prev.Next
			}
		}
	}


	return ret.Next
}
