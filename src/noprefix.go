package main

import "fmt"

type Node struct {
	ch          rune
	children    [10]*Node
	prefixCount int32
	wordCount   int32
}

func NewNode(r rune) *Node {
	return &Node{
		ch:          r,
		children:    [10]*Node{},
		prefixCount: 0,
		wordCount:   0,
	}
}

type Trie struct {
	root *Node
}

func (t *Trie) insert(word string) bool {
	// if t == nil || t.root == nil {
	// 	t = &Trie{NewNode('\000')}
	// }

	curr := t.root
	for _, r := range word {

		index := r - 'a'
		if curr.children[index] == nil {
			curr.children[index] = NewNode(r)
		}

		if curr.children[index].wordCount > 0 {
			return false
		}

		curr.children[index].prefixCount += 1

		curr = curr.children[index]
	}

	if curr.prefixCount > 1 {
		return false
	}

	curr.wordCount += 1

	return true
}

// func (t *Trie) search(word string) bool {
//     if t == nil || t.root == nil {
//         return false
//     }

//     curr := t.root
//     for _, r := range word {
//         index := r - 'a'
//         if curr.children[index] == nil {
//             return false
//         }

//         curr = curr.children[index]
//     }

//     return true
// }

func noPrefix(words []string) {
	// Write your code here

	trie := &Trie{NewNode('\000')}
	for _, word := range words {
		if !trie.insert(word) {
			fmt.Println("BAD SET")
			fmt.Println(word)
			return
		}
	}

	fmt.Println("GOOD SET")

}
