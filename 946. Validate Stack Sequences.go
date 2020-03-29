package main

func validateStackSequences(pushed []int, popped []int) bool {
	s := []int{}
	p1, p2 := 0, 0
	l1, l2 := len(pushed), len(popped)
	if l1 != l2 {
		return false
	}
	for p2 < l2 {
		if len(s) == 0 || s[len(s) - 1] != popped[p2] {
			if p1 == l1 {
				return false
			}
			s = append(s, pushed[p1])
			p1 ++
		} else {
			s = s[:len(s) - 1]
			p2 ++
		}
	}
	if p2 == l2 {
		return true
	} else {
		return false
	}
}
