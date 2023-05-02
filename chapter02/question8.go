package chapter02

import (
	"math"

	"github.com/syhily/code-interviews/common"
)

func shortestArray(target int, nums ...int) int {
	start, end, sum, shortest := 0, 0, nums[0], math.MaxInt

	for start <= end {
		if sum >= target {
			shortest = common.MinNumber(shortest, end-start+1)
			sum -= nums[start]
			start++
		} else {
			end++
			if end >= len(nums) {
				break
			}
			sum += nums[end]
		}
	}

	if shortest == math.MaxInt {
		return 0
	} else {
		return shortest
	}
}
