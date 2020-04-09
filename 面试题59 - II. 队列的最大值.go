package main

type MaxQueue struct {
	q []int
	m []int
}


func Constructor() MaxQueue {
	return MaxQueue{}
}


func (this *MaxQueue) Max_value() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.m[0]
}


func (this *MaxQueue) Push_back(value int)  {
	for len(this.m) > 0 && this.m[len(this.m) - 1] < value {
		this.m = this.m[:len(this.m) - 1]
	}
	this.m = append(this.m, value)
	this.q = append(this.q, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	tmp := this.q[0]
	this.q = this.q[1:]
	if tmp == this.m[0] {
		this.m = this.m[1:]
	}
	return tmp
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
