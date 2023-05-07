package chapter10

type (
	Trie interface {
		insert(word string)
		search(word string) bool
		startWith(prefix string) bool
	}

	node struct {
		word     bool
		children [26]*node
	}

	trie struct {
		root node
	}
)

func newTrie() Trie {
	return &trie{}
}

func (t *trie) insert(word string) {
	curr := &t.root
	for _, r := range word {
		i := r - 'a'
		n := curr.children[i]
		if n == nil {
			n = &node{}
			curr.children[i] = n
		}
		curr = n
	}
	curr.word = true
}

func (t *trie) search(word string) bool {
	curr := &t.root
	for _, r := range word {
		curr = curr.children[r-'a']
		if curr == nil {
			return false
		}
	}
	return curr.word
}

func (t *trie) startWith(prefix string) bool {
	curr := &t.root
	for _, r := range prefix {
		curr = curr.children[r-'a']
		if curr == nil {
			return false
		}
	}
	return true
}
