package main

import "strings"

func stringMatching(words []string) []string {
	ret := []string{}
	for k, s := range words {
		for i := 0; i < len(words); i++ {
			if i == k {
				continue
			}
			if strings.Index(words[i], s) > -1 {
				ret = append(ret, s)
				break
			}
		}
	}
	return ret
}
