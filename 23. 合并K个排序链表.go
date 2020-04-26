package main

import "container/heap"

type Heap []*ListNode

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{})  {
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	x := (*h)[(*h).Len() - 1]
	*h = (*h)[:(*h).Len() - 1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := Heap{}
	heap.Init(&h)
	head := ListNode{}
	p := &head
	for _, list := range lists {
		if list != nil {
			heap.Push(&h, list)
		}
	}

	for h.Len() != 0 {
		tmp := ListNode{}
		cur := heap.Pop(&h).(*ListNode)
		tmp.Val = cur.Val
		p.Next = &tmp
		p = p.Next
		cur = cur.Next
		if cur != nil {
			heap.Push(&h, cur)
		}
	}

	return head.Next

}