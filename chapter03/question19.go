package chapter03

func validateIsPalindrome(s string) bool {
	runes := []rune(s)
	return innerValidateIsPalindrome(runes, 0, len(runes)-1, true)
}

func innerValidateIsPalindrome(runes []rune, begin, end int, canSkip bool) bool {
	for begin < end {
		l := runes[begin]
		r := runes[end]

		if l != r {
			if canSkip {
				return innerValidateIsPalindrome(runes, begin+1, end, false) ||
					innerValidateIsPalindrome(runes, begin, end-1, false)
			} else {
				return false
			}
		}

		begin++
		end--
	}

	return true
}
