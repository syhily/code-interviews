package chapter02

func multipleNumberBelowK(target int, nums ...int) int {
	start, multiple, counts := 0, 1, 0

	for end, num := range nums {
		multiple *= num

		for multiple >= target && start <= end {
			multiple /= nums[start]
			start++
		}

		counts += end - start + 1
	}

	return counts
}
