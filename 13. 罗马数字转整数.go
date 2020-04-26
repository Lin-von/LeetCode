package main

func romanToInt(s string) int {
	v := []int {
		73 : 1,
		86 : 5,
		88 : 10,
		76 : 50,
		67 : 100,
		68 : 500,
		77 : 1000,
	}

	n := len(s)
	sum := v[s[n - 1]]
	for i := 0; i < n - 1; i++ {
		if v[s[i]] < v[s[i + 1]] {
			sum -= v[s[i]]
		} else {
			sum += v[s[i]]
		}
	}

	return sum
}

