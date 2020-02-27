package main

import (
	"container/list"
)

type LRUNode struct {
	Key int
	Val int
}

type LRUCache struct {
	cap int
	l *list.List
	hash map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:  capacity,
		l:    new(list.List),
		hash: map[int]*list.Element{},
	}
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.hash[key]; ok {
		this.l.MoveToFront(this.hash[key])
		return this.hash[key].Value.(LRUNode).Val
	} else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.hash[key]; ok {
		this.l.MoveToFront(node)
		node.Value = LRUNode{
			Key: key,
			Val: value,
		}
	} else {
		tmp := this.l.PushFront(LRUNode{
			Key: key,
			Val: value,
		})
		this.hash[key] = tmp

		if this.l.Len() > this.cap {
			v := this.l.Remove(this.l.Back())
			delete(this.hash, v.(LRUNode).Key)
		}
	}
}
