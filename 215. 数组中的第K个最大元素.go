package main

var arr []int

func findK(l, r, k int) int {
	if l == r {
		return arr[l]
	}

	pos := l
	for i := l + 1; i <= r; i++ {
		if arr[l] < arr[i] {
			pos ++
			arr[i], arr[pos] = arr[pos], arr[i]
		}
	}
	arr[l], arr[pos] = arr[pos], arr[l]
	if pos == k {
		return arr[pos]
	} else if pos < k {
		return findK(pos + 1, r, k)
	} else {
		return findK(l, pos - 1, k)
	}
}

func findKthLargest(nums []int, k int) int {
	arr = nums
	return findK(0, len(nums) - 1, k - 1)
}
