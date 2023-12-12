package main

func caesarCipher(s string, k int32) string {
	// Write your code here
	res := []rune{}

	for _, c := range s {
		res = append(res, c+k)
	}

	return string(res)
}
