package main

func intToRoman(num int) string {
	m := map[int]string {
		1000:"M",
		900:"CM",
		500:"D",
		400:"CD",
		100:"C",
		90:"XC",
		50:"L",
		40:"XL",
		10:"X",
		9:"IX",
		5:"V",
		4:"IV",
		1:"I",
	}

	n := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	pos := 0
	ret := ""
	for num > 0 {
		if num < n[pos] {
			pos ++
		} else {
			num -= n[pos]
			ret += m[n[pos]]
		}
	}
	return ret
}
