package main

func nextNum(n int) int {
	ret := 0
	for n != 0 {
		tmp := n % 10
		ret += tmp * tmp
		n /= 10
	}
	return ret
}

func isHappy(n int) bool {
	slow := n
	fast := nextNum(n)
	for fast != 1 && slow != fast {
		fast = nextNum(nextNum(fast))
		slow = nextNum(slow)
	}
	return fast == 1
}
