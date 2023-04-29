package chapter03

func findAllStartIndexOfInclusionWords(s1, s2 string) []int {
	var res []int

	// Count all the runes in s2.
	counts := make([]int, 26)
	for _, r := range s2 {
		counts[r-'a']++
	}

	l := len(s2)
	for i, r := range s1 {
		counts[r-'a']--
		if i >= l {
			counts[s1[i-l]-'a']++
		}

		if i >= l-1 && allValueIsZero(counts) {
			res = append(res, i-len(s2)+1)
		}
	}

	return res
}
