package main

func part(l, r int, arr []int) {
	if l >= r {
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
	part(pos + 1, r, arr)
	part(l, pos - 1, arr)
}

func sortArray(nums []int) []int {
	part(0, len(nums) - 1, nums)
	return nums
}
