package chapter05

func isSortedWords(alphabet string, words ...string) bool {
	a := []rune(alphabet)
	idx := make(map[rune]int, len(alphabet))

	for i := 0; i < len(alphabet); i++ {
		idx[a[i]] = i
	}

	left := []rune(words[0])
	for i := 1; i < len(words); i++ {
		right := []rune(words[i])

		if !isSortedRuneSlice(left, right, idx) {
			return false
		}

		left = right
	}

	return true
}

func isSortedRuneSlice(left, right []rune, idx map[rune]int) bool {
	for j := 0; j < len(left) && j < len(right); j++ {
		l, r := idx[left[j]], idx[right[j]]
		if l > r {
			return false
		}

		if r > l {
			return true
		}
	}

	return len(left) <= len(right)
}
