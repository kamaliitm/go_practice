package main

func reverseWordsInSentence(sent string) string {

	sent = reverseStringV2(sent, 0, len(sent)-1)
	i := 0
	start := 0
	for i < len(sent) {
		start = i
		for i < len(sent) {
			if sent[i] == ' ' {
				sent = reverseStringV2(sent, start, i-1)
				i++
				break
			}
			i++
		}
	}

	sent = reverseStringV2(sent, start, i-1)

	return sent
}

func reverseStringV2(str string, startInd int, endInd int) string {
	runes := []rune(str)
	for i, j := startInd, endInd; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
