package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	ori := head

	if head == nil || head.Next == nil {
		return nil
	}

	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}
	slow.Next = slow.Next.Next
	return ori
}
