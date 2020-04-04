package main

import "strconv"

type NTrie struct {
	isNum bool
	Next [10]*RTrie
}

func dfs(cur, ret string, num int)  {

}

func minNumber(nums []int) string {
	t := NTrie{}
	for _, num := range nums {
		s := strconv.Itoa(num)
		tmp := t
		for i := 0; i < len(s); i++ {
			if tmp.Next[s[i] - '0'] == nil {

			}
		}
	}
}
