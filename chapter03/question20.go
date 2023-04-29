package chapter03

func countAllPalindrome(s string) int {
	counts := 0

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		counts += innerCountAllPalindrome(runes, i, i)
		counts += innerCountAllPalindrome(runes, i, i+1)
	}

	return counts
}

func innerCountAllPalindrome(runes []rune, i int, j int) int {
	counts := 0

	for i >= 0 && j < len(runes) {
		if runes[i] == runes[j] {
			counts++
			i--
			j++
		} else {
			break
		}
	}

	return counts
}
