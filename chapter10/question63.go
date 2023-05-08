package chapter10

import "strings"

const space = " "

func replaceWords(sentence string, dict ...string) string {
	t := &trie{}
	for _, d := range dict {
		t.insert(d)
	}

	words := strings.Split(sentence, space)

	for i, word := range words {
		words[i] = replaceWord(t, word)
	}

	return strings.Join(words, space)
}

func replaceWord(t *trie, word string) string {
	curr := &t.root
	for i, w := range word {
		n := curr.children[w-'a']
		if n == nil {
			return word
		}

		if n.isWord {
			return word[:i+1]
		}

		curr = n
	}

	return word
}
