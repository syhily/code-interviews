package chapter03

func findTheShortestSubstring(s, t string) string {
	// Count all the runes in t.
	runes1, runes2, counts := []rune(t), []rune(s), make(map[rune]int)
	for _, r := range runes1 {
		counts[r]++
	}

	// Find the min start and length.
	short, length, start, end := 0, 0, 0, 0
	for start < len(s) {
		r := runes2[start]
		if _, ok := counts[r]; !ok {
			start++
			continue
		}

		if end < start {
			end = start
		}

		for end < len(s) {
			r2 := runes2[end]
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

	return string(runes2[short : short+length])
}

func allValueIsBelowZero(counts map[rune]int) bool {
	for _, v := range counts {
		if v > 0 {
			return false
		}
	}
	return true
}
