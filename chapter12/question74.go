package chapter12

import "github.com/syhily/code-interviews/common"

func mergeIntervals(intervals ...[]int) [][]int {
	// Bubble sort the intervals
	for i := range intervals {
		for j := 0; j < len(intervals)-i-1; j++ {
			l := intervals[j]
			r := intervals[j+1]

			if l[0] > r[0] {
				intervals[j], intervals[j+1] = r, l
			}
		}
	}

	var res [][]int

	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
			continue
		}

		l := len(res) - 1
		tail := res[l]

		if interval[0] <= tail[1] {
			tail[1] = common.MaxNumber(tail[1], interval[1])
		} else {
			res = append(res, interval)
		}
	}

	return res
}
