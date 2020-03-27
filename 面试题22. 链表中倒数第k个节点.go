package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	tmp := head
	for k != 0 {
		tmp = tmp.Next
		k --
	}
	for tmp != nil {
		tmp = tmp.Next
		head = head.Next
	}
	return head
}
