package main

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	found := false
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if !found {
				continue
			}
			return cnt
		} else {
			found = true
			cnt ++
		}
	}
	return cnt
}