package main

var dp []int
var B int

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func dfs(k, n int) int {
	if n == 0 {
		return 0
	}

	if k == 1 {
		dp[(k - 1) * B + n] = n
	}

	if n == 1 {
		dp[(k - 1) * B + n] = 1
	}

	if dp[(k - 1) * B + n] != 0 {
		return dp[(k - 1) * B + n]
	}

	l := 1
	r := n
	for l + 1 < r {
		x := (l + r) >> 1
		t1 := dfs(k - 1, x - 1)
		t2 := dfs(k, n - x)

		if t1 < t2 {
			l = x
		} else if t1 > t2 {
			r = x
		} else {
			l = x
			r = x
		}
	}
	dp[(k - 1) * B + n] = min(max(dfs(k - 1, l - 1), dfs(k, n - l)),
		max(dfs(k - 1, r - 1), dfs(k, n - r))) + 1
	return dp[(k - 1) * B + n]
}

func superEggDrop(K int, N int) int {
	dp = make([]int, K * (N + 1))
	B = N + 1
	return dfs(K, N)
}
