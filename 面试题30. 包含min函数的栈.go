package main

type MinStack struct {
	stack []int
	minV []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	s := MinStack{}
	return s
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	clen := len(this.minV)
	if clen == 0 || x < this.minV[clen - 1] {
		this.minV = append(this.minV, x)
	} else {
		this.minV = append(this.minV, this.minV[clen - 1])
	}
}


func (this *MinStack) Pop()  {
	if len(this.stack) == 0 {
		return
	}

	this.stack = this.stack[0:len(this.stack) - 1]
	this.minV = this.minV[0:len(this.minV) - 1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) Min() int {
	return this.minV[len(this.minV) - 1]
}



