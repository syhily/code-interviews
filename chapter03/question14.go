package chapter03

func findInclusionWord(s1, s2 string) bool {
	// Count all the runes in s1.
	counts := make([]int, 26)
	for _, r := range s1 {
		counts[r-'a']++
	}

	l := len(s1)
	for i, r := range s2 {
		counts[r-'a']--
		if i >= l {
			counts[s2[i-l]-'a']++
		}

		if i >= l-1 && allValueIsZero(counts) {
			return true
		}
	}

	return false
}

func allValueIsZero(c []int) bool {
	for _, v := range c {
		if v != 0 {
			return false
		}
	}
	return true
}
