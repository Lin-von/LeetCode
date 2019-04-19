/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode
type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l ListNode
	c := 0
	l.Val = l1.Val + l2.Val
	c = l.Val / 10
	l.Val = l.Val % 10

	p := &l
	l1 = l1.Next
	l2 = l2.Next

	for l1 != nil && l2 != nil {
		p.Next = new(ListNode)
		p = p.Next

		p.Val = l1.Val + l2.Val + c
		c = p.Val / 10
		p.Val = p.Val % 10


		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		p.Next = new(ListNode)
		p = p.Next

		p.Val = l1.Val + c
		c = p.Val / 10
		p.Val = p.Val % 10


		l1 = l1.Next

	}

	for l2 != nil {
		p.Next = new(ListNode)
		p = p.Next

		p.Val = l2.Val + c
		c = p.Val / 10
		p.Val = p.Val % 10


		l2 = l2.Next

	}

	if c > 0 {
		p.Next = new(ListNode)
		p = p.Next

		p.Val = c
	}

	return &l
}