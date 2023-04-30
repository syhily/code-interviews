package chapter12

import "github.com/syhily/code-interviews/common"

func mergeIntervals(intervals ...[]int) [][]int {
	// Bubble sort the intervals
	for i := range intervals {
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][0] > intervals[j+1][0] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}

	var res [][]int

	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
			continue
		}

		tail := res[len(res)-1]

		if interval[0] <= tail[1] {
			tail[1] = common.MaxNumber(tail[1], interval[1])
		} else {
			res = append(res, interval)
		}
	}

	return res
}
