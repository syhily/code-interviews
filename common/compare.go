package common

func MaxNumber[N Number](a, b N) N {
	if a > b {
		return a
	}
	return b
}

func MinNumber[N Number](a, b N) N {
	if a < b {
		return a
	}
	return b
}
