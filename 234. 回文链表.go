package main

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	ret := true

	slow := head
	next := slow.Next
	fast := head.Next

	head.Next = nil
	for fast != nil && fast.Next != nil {
		// 快指针先移动
		fast = fast.Next.Next
		// 慢指针移动，同时反转链表
		tmp := next.Next
		next.Next = slow
		slow = next
		next = tmp
	}

	// 此时快指针遍历到尾部，慢指针在链表中间或者左链表尾部
	p := slow
	q := next
	// 如果链表长度为奇数,移动一下跳过中间位置
	if fast == nil {
		p = p.Next
	}

	// 比较链表
	for p != nil {
		if p.Val != q.Val {
			ret = false
		}
		p = p.Next
		q = q.Next
	}

	// 恢复被翻转的链表
	pre := slow.Next
	slow.Next = next
	for pre != nil {
		tmp := pre.Next
		pre.Next = slow
		slow = pre
		pre = tmp
	}

	return ret
}