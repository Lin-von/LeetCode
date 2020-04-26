package main

func hammingWeight(num uint32) int {
	ret := 0
	for num != 0 {
		if num & 1 == 1 {
			ret ++
		}
		num >>= 1
	}
	return ret
}
