package main

import "strconv"

func multiply(num1 string, num2 string) string {
	res := make([]int, len(num1) + len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			res[i + j + 1] += n1 * n2
			res[i + j] += res[i + j + 1] / 10
			res[i + j + 1] %= 10
		}
	}

	pos := 0
	for pos < len(res) && res[pos] == 0 {
		pos ++
	}

	if pos == len(res) {
		return "0"
	}

	ret := ""
	for pos < len(res) {
		ret += strconv.Itoa(res[pos])
		pos ++
	}


	return ret
}
