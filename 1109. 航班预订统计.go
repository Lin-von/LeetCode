package main

func corpFlightBookings(bookings [][]int, n int) []int {
	cnt := make([]int, n + 1)
	ret := []int{}
	for _, c := range bookings {
		cnt[c[0] - 1] += c[2]
		cnt[c[1]] -= c[2]
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += cnt[i]
		ret = append(ret, sum)
	}
	return ret
}
