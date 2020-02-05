package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB
	swa, swb := false, false

	if pa == nil || pb == nil {
		return nil
	}

	for pa != pb {
		pa = pa.Next
		pb = pb.Next

		if pa == nil {
			if !swa {
				pa = headB
				swa = true
			} else {
				return nil
			}
		}

		if pb == nil {
			if !swb{
				pb = headA
				swb = true
			} else {
				return nil
			}
		}
	}
	return pa
}
