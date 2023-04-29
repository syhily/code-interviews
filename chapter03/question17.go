package chapter03

func findTheShortestSubstring(s, t string) string {
	// Count all the runes in t.
	counts := make(map[rune]int)
	for _, r := range t {
		counts[r]++
	}

	// Find the min start and length.
	short, length, start, end := 0, 0, 0, 0
	for start < len(s) {
		r := rune(s[start])
		if _, ok := counts[r]; !ok {
			start++
			continue
		}

		if end < start {
			end = start
		}

		for end < len(s) {
			r2 := rune(s[end])
			if c, ok := counts[r2]; ok {
				counts[r2] = c - 1
			}

			if allValueIsBelowZero(counts) {
				if length == 0 || end-start+1 < length {
					length = end - start + 1
					short = start
				}
				break
			}
			end++
		}

		counts[r]++
		start++
	}

	return s[short : short+length]
}

func allValueIsBelowZero(counts map[rune]int) bool {
	for _, v := range counts {
		if v > 0 {
			return false
		}
	}
	return true
}
