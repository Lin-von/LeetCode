package main

//func generateParenthesis(n int) []string {
//	ret := []string{}
//	if n == 0 {
//		ret = append(ret, "")
//		return ret
//	}
//
//	for i := 0; i < n; i++ {
//		for _, l := range generateParenthesis(i) {
//			for _, r := range generateParenthesis(n - 1 - i) {
//				ret = append(ret, "(" + l + ")" + r)
//			}
//		}
//	}
//	return ret
//}


func Scan(arr *[]string, l int, r int, n int, s string) {
	if l + r == n * 2 {
		*arr = append(*arr, s)
		return
	}

	if l < n {
		Scan(arr, l + 1, r, n, s + "(")
	}

	if l > r {
		Scan(arr, l, r + 1, n, s + ")")
	}
}

func generateParenthesis(n int) []string {
	ret := []string{}
	Scan(&ret, 0, 0, n, "")
	return ret
}
