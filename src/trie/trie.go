package trie

type Trie struct {
	children  []*Trie
	isWordEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	if this == nil {
		return
	}

	curr := this
	var parent *Trie
	for _, r := range word {
		if curr.children == nil {
			curr.children = make([]*Trie, 26)
		}

		index := r - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &Trie{}
		}

		parent = curr
		curr = curr.children[index]
	}

	parent.isWordEnd = true
}

func (this *Trie) Search(word string) bool {
	if this == nil {
		return false
	}

	curr := this
	var parent *Trie
	for _, r := range word {
		index := r - 'a'
		if curr.children == nil || curr.children[index] == nil {
			return false
		}

		parent = curr
		curr = curr.children[index]
	}

	return parent != nil && parent.isWordEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	if this == nil {
		return false
	}

	curr := this
	for _, r := range prefix {
		index := r - 'a'
		if curr.children == nil || curr.children[index] == nil {
			return false
		}

		curr = curr.children[index]
	}

	return true
}
