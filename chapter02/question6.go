package chapter02

func twoSum(sum int, arrays ...int) (l int, r int) {
	l, r = 0, len(arrays)-1

	for l < r {
		curr := arrays[l] + arrays[r]

		if curr > sum {
			r--
		} else if curr < sum {
			l++
		} else {
			break
		}
	}

	return
}
