package chapter03

func findAllStartIndexOfInclusionWords(s1, s2 string) []int {
	var res []int

	// Count all the runes in s2.
	counts := make(map[rune]int, 26)
	for _, r := range []rune(s2) {
		counts[r]++
	}

	runes := []rune(s1)
	l := len(s2)
	for i, r := range runes {
		counts[r]--
		if i >= l {
			counts[runes[i-l]]++
		}

		if i >= l-1 && allValueIsZero(counts) {
			res = append(res, i-len(s2)+1)
		}
	}

	return res
}
