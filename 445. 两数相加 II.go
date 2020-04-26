package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := []int{}
	s2 := []int{}

	c := 0
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	var ret *ListNode = nil
	p1 := len(s1) - 1
	p2 := len(s2) - 1
	for p1 >= 0 && p2 >= 0 {
		n := ListNode{}
		tmp := s1[p1] + s2[p2]
		n.Val = (c + tmp) % 10
		c = (c + tmp) / 10
		n.Next = ret
		ret = &n
		p1 --
		p2 --
	}

	for p1 >= 0 {
		n := ListNode{}
		n.Val = (c + s1[p1]) % 10
		c = (c + s1[p1]) / 10
		n.Next = ret
		ret = &n
		p1 --
	}

	for p2 >= 0 {
		n := ListNode{}
		n.Val = (c + s2[p2]) % 10
		c = (c + s2[p2]) / 10
		n.Next = ret
		ret = &n
		p2 --
	}

	if c != 0 {
		n := ListNode{}
		n.Val = c
		n.Next = ret
		ret = &n
	}


	return ret

}
