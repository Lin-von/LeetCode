package main

func numOfWays(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 12
	} else {
		var tw, th int64
		tw = 6
		th = 6

		for i := 1; i < n; i++ {
			t1 := (tw * 3) % (1e9 + 7) + (th * 2) % (1e9 + 7)
			t2 := (tw * 2) % (1e9 + 7) + (th * 2) % (1e9 + 7)
			tw = t1
			th = t2
		}
		return (int)(tw + th) % (1e9 + 7)
	}
}
