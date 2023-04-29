package chapter03

func findInclusionWord(s1, s2 string) bool {
	// Count all the runes in s1.
	counts := make(map[rune]int, 26)
	for _, r := range []rune(s1) {
		counts[r]++
	}

	runes := []rune(s2)
	l := len(s1)
	for i, r := range runes {
		counts[r]--
		if i >= l {
			counts[runes[i-l]]++
		}

		if i >= l-1 && allValueIsZero(counts) {
			return true
		}
	}

	return false
}

func allValueIsZero(c map[rune]int) bool {
	for _, v := range c {
		if v != 0 {
			return false
		}
	}
	return true
}
