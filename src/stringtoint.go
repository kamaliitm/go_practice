package main

import "fmt"

func intToString(n int32) string {
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	res := ""
	for n > 0 {
		d := n % 10
		res = fmt.Sprintf("%v%s", d, res)
		n /= 10
	}

	if isNegative {
		res = fmt.Sprintf("-%s", res)
	}

	return res
}

func stringToInt(s string) int32 {
	if len(s) <= 0 {
		return 0
	}

	var res int32 = 0
	isNegative := false
	if s[0] == '-' {
		isNegative = true
		s = s[1:]
	}

	for _, chr := range s {
		d := chr - '0'
		res = (res * 10) + d
	}

	if isNegative {
		res = -res
	}

	return res
}
