package main

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}
	s := []*ListNode{}
	p1 := head
	p2 := head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	for p1.Next != nil {
		tmp := p1.Next
		s = append(s, tmp)
		p1.Next = nil
		p1 = tmp
	}

	p := head
	for p != nil {
		next := p.Next
		if len(s) == 0 {
			break
		}
		p.Next = s[len(s) - 1]
		p = p.Next
		p.Next = next
		s = s[:len(s) - 1]
		p = p.Next
	}
}
