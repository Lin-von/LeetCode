package main

//翻转数字的一半比较大小即可
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0){
		return false
	}
	t := x
	r := 0
	for ; t != 0 && t > r; {
		r = r * 10 + t % 10
		t /= 10
	}
	if t == r || t == r/10 {
		return true
	} else {
		return false
	}
}