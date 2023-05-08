package chapter10

func shortestEncodingLength(words ...string) int {
	root := &node{}
	for _, w := range words {
		curr := root
		for i := len(w) - 1; i >= 0; i-- {
			u := w[i] - 'a'
			if curr.children[u] == nil {
				curr.children[u] = &node{}
			}
			curr = curr.children[u]
		}
		curr.isWord = true
	}

	// Calculate the total length.
	sum := 0
	countWords(root, 1, &sum)
	return sum
}

func countWords(node *node, depth int, sum *int) {
	isLeaf := true
	for _, child := range node.children {
		if child != nil {
			isLeaf = false
			countWords(child, depth+1, sum)
		}
	}
	if isLeaf {
		*sum += depth
	}
}
