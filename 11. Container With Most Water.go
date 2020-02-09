package main

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max := 0
	for l < r {
		area := r - l
		if height[r] > height[l] {
			area = area * height[l]
			l++
		} else {
			area = area * height[r]
			r--
		}

		if area > max {
			max = area
		}

	}
	return max
}
