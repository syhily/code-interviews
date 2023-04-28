package chapter3

import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)
	begin, end := 0, len(runes)-1

	for begin < end {
		for !unicode.IsLetter(runes[begin]) && begin <= end {
			begin++
		}
		l := unicode.ToLower(runes[begin])

		for !unicode.IsLetter(runes[end]) && begin <= end {
			end--
		}
		r := unicode.ToLower(runes[end])

		if l != r {
			return false
		}

		begin++
		end--
	}

	return true
}
