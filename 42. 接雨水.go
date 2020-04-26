package main

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	p, q := 0, 0
	l, r := 0, len(height) - 1
	sum := 0

	for l < r {
		if height[l] < height[r] {
			if height[l] >= p {
				p = height[l]
			} else {
				sum += p - height[l]
			}
			l ++
		} else {
			if height[r] >= q {
				q = height[r]
			} else {
				sum += q - height[r]
			}
			r --
		}
	}

	return sum
}
