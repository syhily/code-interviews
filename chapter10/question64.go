package chapter10

type MagicDictionary interface {
	buildDict(words []string)
	search(word string) bool
}

type magicDictionary struct {
	root node
}

func (m *magicDictionary) buildDict(words []string) {
	for _, word := range words {
		curr := &m.root
		for _, w := range word {
			i := w - 'a'
			n := curr.children[i]
			if n == nil {
				n = &node{}
				curr.children[i] = n
			}
			curr = n
		}
		curr.isWord = true
	}
}

func (m *magicDictionary) search(word string) bool {
	return innerSearch(&m.root, word, false)
}

func innerSearch(curr *node, word string, edited bool) bool {
	for i, w := range word {
		if curr.children[w-'a'] != nil {
			curr = curr.children[w-'a']
			continue
		}

		if edited {
			return false
		}

		for _, child := range curr.children {
			if child != nil && innerSearch(child, word[i+1:], true) {
				return true
			}
		}
	}

	return curr.isWord && edited
}
