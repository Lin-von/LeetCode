package main

func twoSum(n int) []float64 {

	if n == 1 {
		ret := make([]float64, 6)
		for i := 0; i < 6; i++ {
			ret[i] = 1.0 / 6
		}
		return ret
	}

	tmp := twoSum(n - 1)
	ret := make([]float64, len(tmp) + 5)
	for i := 0; i < 6; i++ {
		for j := 0; j < len(tmp); j++ {
			ret[j + i] += tmp[j] / 6
		}
	}
	return ret
}
