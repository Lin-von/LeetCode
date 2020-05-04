package main

func plusOne(digits []int) []int {
	n := len(digits) - 1
	for n >= 0 && digits[n] == 9 {
		digits[n] = 0
		n --
	}
	if n < 0 {
		digits = append([]int{1}, digits...)
		return digits
	} else {
		digits[n] ++
		return digits
	}
}
