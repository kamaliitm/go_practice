package main

func baseConvert(str string, b1 int32, b2 int32) string {
	if len(str) == 0 {
		return ""
	}

	isNegative := false
	if str[0] == '-' {
		isNegative = true
	}

	var numBase10 int32 = 0
	for _, c := range str {
		if isDigit(c) {
			numBase10 = (numBase10 * b1) + (c - '0')
		} else {
			numBase10 = (numBase10 * b1) + (c - 'A' + 10)
		}
	}

	runes := []rune{}
	for numBase10 > 0 {
		remainder := numBase10 % b2
		if remainder >= 10 {
			runes = append(runes, 'A'+remainder-10)
		} else {
			runes = append(runes, '0'+remainder)
		}

		numBase10 /= b2
	}

	if isNegative {
		runes = append(runes, '-')
	}

	strBaseB2 := reverseString(runes)

	return strBaseB2
}

func reverseString(runes []rune) string {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[j], runes[i] = runes[i], runes[j]
	}

	return string(runes)
}

func isDigit(c rune) bool {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, digit := range digits {
		if c == digit {
			return true
		}
	}

	return false
}
