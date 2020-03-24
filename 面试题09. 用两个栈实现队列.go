package main


type CQueue struct {
	stack1 []int
	stack2 []int
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.stack1 = append(this.stack1, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return -1
	}
	if len(this.stack2) > 0 {
		ret := this.stack2[len(this.stack2) - 1]
		this.stack2 = this.stack2[:len(this.stack2) - 1]
		return ret
	} else {
		for i := len(this.stack1) - 1; i > 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		ret := this.stack1[0]
		this.stack1 = this.stack1[:0]
		return ret
	}
}

