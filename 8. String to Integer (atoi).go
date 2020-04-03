package main

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	pos := 0
	for pos < len(str) && str[pos] == ' ' {
		pos ++
	}
	if pos == len(str) {
		return 0
	}
	var f int64 = 0
	if str[pos] == '+' {
		pos ++
	} else if str[pos] == '-' {
		f = 1
		pos ++
	}
	var ret int64
	ret = 0
	for pos < len(str) && str[pos] >= '0' && str[pos] <= '9' {
		ret *= 10
		ret += int64(str[pos] - '0')
		if ret - int64(math.MaxInt32) - f > 0 {
			if f == 1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		pos ++
	}
	res := int(ret)
	if f == 1 {
		return res * -1
	} else {
		return res
	}

}
