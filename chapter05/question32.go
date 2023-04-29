package chapter05

// s1, s2 should be a word which contains only lower case letters.
func isAnagramWithLowerCaseLetters(s1, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)

	l := len(r1)
	if l != len(r2) || s1 == s2 {
		return false
	}

	counts := make([]int, 26)
	for i := 0; i < l; i++ {
		counts[r1[i]-'a']++
	}

	for i := 0; i < l; i++ {
		idx := r2[i] - 'a'
		c := counts[idx]
		if c == 0 {
			return false
		}
		counts[idx] = c - 1
	}

	return true
}

func isAnagram(s1, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)

	l := len(r1)
	if l != len(r2) || s1 == s2 {
		return false
	}

	counts := make(map[rune]int)
	for i := 0; i < l; i++ {
		counts[r1[i]]++
	}

	for i := 0; i < l; i++ {
		c := counts[r2[i]]
		if c == 0 {
			return false
		}
	}

	return true
}
