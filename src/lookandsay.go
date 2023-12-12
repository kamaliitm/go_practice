package main

import "fmt"

func lookAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		s = nextString(s)
	}

	return s
}

func nextString(s string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		currCount := 1
		j := i + 1
		for j < len(s) && s[j] == s[i] {
			currCount++
			j++
			i++
		}

		ret += fmt.Sprintf("%v%s", currCount, string(s[i]))
	}

	return ret
}
