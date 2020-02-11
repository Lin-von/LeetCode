package main

import "container/heap"

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{})  {
	*h = append(*h, x.(int))
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
		for list != nil {
			heap.Push(&h, list.Val)
			list = list.Next
		}
	}

	for h.Len() != 0 {
		tmp := ListNode{}
		tmp.Val = heap.Pop(&h).(int)
		p.Next = &tmp
		p = p.Next
	}

	return head.Next

}