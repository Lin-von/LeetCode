package main

var gS string

var sh = map[string][]string{
	"2" : []string{"a", "b", "c"},
	"3" : []string{"d", "e", "f"},
	"4" : []string{"g", "h", "i"},
	"5" : []string{"j", "k", "l"},
	"6" : []string{"m", "n", "o"},
	"7" : []string{"p", "q", "r", "s"},
	"8" : []string{"t", "u", "v"},
	"9" : []string{"w", "x", "y", "z"},
}

func scan(i int, arr *[]string, s string)  {
	if i == len(gS) {
		if len(s) > 0 {
			*arr = append(*arr, s)
		}
		return
	}

	for _, a := range sh[string(gS[i])] {
		scan(i + 1, arr, s + a)
	}
}

func letterCombinations(digits string) []string {
	gS = digits
	ret := []string{}
	scan(0, &ret, "")
	return ret
}
