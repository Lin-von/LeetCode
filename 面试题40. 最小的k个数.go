package main

func partK(l, r, k int, arr []int) {
	if l == r {
		return
	}

	pos := l
	for i := l + 1; i <= r; i++ {
		if arr[l] > arr[i] {
			pos ++
			arr[i], arr[pos] = arr[pos], arr[i]
		}
	}
	arr[l], arr[pos] = arr[pos], arr[l]
	if pos == k {
		return
	} else if pos < k {
		partK(pos + 1, r, k, arr)
	} else {
		partK(l, pos - 1, k, arr)
	}
}

func getLeastNumbers(arr []int, k int) []int {
	partK(0, len(arr) - 1, k, arr)
	return arr[0:k]
}
