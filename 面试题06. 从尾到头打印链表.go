package main

func reversePrint(head *ListNode) []int {
	ret := []int{}
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	p := 0
	q := len(ret) - 1
	for p < q {
		ret[p], ret[q] = ret[q], ret[p]
		p ++
		q --
	}
	return ret
}
