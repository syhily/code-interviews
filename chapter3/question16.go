package chapter3

// We can still solve this question by using two pointers.
func lengthOfTheLongestSubString(s string) int {
	start, end, longest := 0, 0, 0
	counts := make(map[rune]int)
	runes := []rune(s)

	for end < len(s) {
		counts[runes[end]]++
		for counts[runes[end]] > 1 {
			counts[runes[start]]--
			start = start + 1
		}

		curr := end - start + 1
		if curr > longest {
			longest = curr
		}

		end++
	}

	return longest
}
