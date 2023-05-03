package chapter02

import "github.com/syhily/code-interviews/common"

func findTheMaxLength(nums ...int) int {
	counts := map[int]int{0: -1}
	sum, max := 0, 0

	for i, num := range nums {
		if num == 0 {
			sum--
		} else {
			sum++
		}

		if j, ok := counts[sum]; ok {
			max = common.MaxNumber(max, i-j)
		} else {
			counts[sum] = i
		}
	}

	return max
}
