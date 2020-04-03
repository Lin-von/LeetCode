package main

import (
"container/heap"
)

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

func (h *Heap) Top() interface{} {
	x := (*h)[0]
	return x
}

type minHeap struct {
	Heap
}

type maxHeap struct {
	Heap
}

func (h maxHeap) Less(i, j int) bool {
	return h.Heap[i] > h.Heap[j]
}

type MedianFinder struct {
	ma maxHeap
	mi minHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{}
	mf.ma.Heap = Heap{}
	heap.Init(&(mf.ma))
	mf.mi.Heap = Heap{}
	heap.Init(&(mf.mi))
	return mf
}


func (this *MedianFinder) AddNum(num int)  {
	if this.ma.Len() == this.mi.Len() {
		heap.Push(&(this.ma), num)
		heap.Push(&(this.mi), heap.Pop(&(this.ma)))
	} else {
		heap.Push(&(this.mi), num)
		heap.Push(&(this.ma), heap.Pop(&(this.mi)))
	}

}


func (this *MedianFinder) FindMedian() float64 {
	if this.ma.Len() == this.mi.Len() {
		a := this.ma.Top().(int)
		b := this.mi.Top().(int)
		return (float64(a + b)) / 2
	} else {
		return float64(this.mi.Top().(int))
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
