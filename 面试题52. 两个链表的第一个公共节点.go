package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB  == nil {
		return nil
	}
	p := headA
	q := headB
	pf := true
	qf := true

	for {
		if p == q {
			return p
		}

		p = p.Next
		if p == nil {
			if pf {
				p = headB
				pf = false
			} else {
				return nil
			}
		}

		q = q.Next
		if q == nil {
			if qf {
				q = headA
				qf = false
			} else {
				return nil
			}
		}
	}
}
