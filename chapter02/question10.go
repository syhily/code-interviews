package chapter02

func sumEqualKArrays(k int, nums ...int) int {
	summary, sum, counts := make(map[int]int), 0, 0
	summary[0] = 1

	for _, num := range nums {
		sum += num
		counts += summary[sum-k]
		summary[sum]++
	}

	return counts
}
