package main

type LadderQueueNode struct {
	word  string
	level int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue := []LadderQueueNode{}
	queue = append(queue, LadderQueueNode{word: beginWord, level: 1})
	queueIter := 0
	visited := map[string]bool{}

	wordDict := getMapFromList(wordList)

	for queueIter < len(queue) {
		node := queue[queueIter]
		if _, ok := visited[node.word]; !ok {
			visited[node.word] = true

			// get all the edges for this word
			runes := []rune(node.word)
			for i, r := range runes {
				for j := 0; j < 26; j++ {
					nextCh := rune('a' + j)
					if nextCh == runes[i] {
						continue
					}

					runes[i] = nextCh
					nextWord := string(runes)

					_, ok1 := wordDict[nextWord]
					_, ok2 := visited[nextWord]

					if ok1 {
						if nextWord == endWord {
							return node.level + 1
						}

						if !ok2 {
							queue = append(queue, LadderQueueNode{word: nextWord, level: node.level + 1})
						}
					}

				}

				runes[i] = r
			}

		}

		queueIter += 1
	}

	return 0
}

func getMapFromList(wordList []string) map[string]bool {
	wordDict := map[string]bool{}

	for _, w := range wordList {
		wordDict[w] = true
	}

	return wordDict
}
