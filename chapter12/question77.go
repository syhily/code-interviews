package chapter12

import "github.com/syhily/code-interviews/common"

func mergeSort(numbers ...int) []int {
	length, res, merge := len(numbers), numbers, make([]int, len(numbers))

	for step := 1; step < length; step *= 2 {
		for start := 0; start < length; start += step * 2 {
			mid := common.MinNumber(start+step, length)
			end := common.MinNumber(mid+step, length)
			l, r, i := start, mid, start

			for l < mid || r < end {
				if r == length || (l < mid && res[l] < res[r]) {
					merge[i] = res[l]
					l++
				} else {
					merge[i] = res[r]
					r++
				}
				i++
			}
		}

		res, merge = merge, res
	}

	return res
}
