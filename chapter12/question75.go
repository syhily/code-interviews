package chapter12

import (
	"math"

	"github.com/syhily/code-interviews/common"
)

func countingSort(numbers ...int) []int {
	max, min := math.MinInt, math.MaxInt
	for _, number := range numbers {
		max = common.MaxNumber(max, number)
		min = common.MinNumber(min, number)
	}

	counts := make([]int, max-min+1)
	for _, number := range numbers {
		counts[number-min]++
	}

	res := make([]int, 0, len(numbers))
	for d, count := range counts {
		for i := 0; i < count; i++ {
			res = append(res, d+min)
		}
	}

	return res
}

func sortByGivenIndex(arr1, arr2 []int) []int {
	// Count the occur times of items in arr1.
	counts := make([]int, 1001)
	for _, i := range arr1 {
		counts[i]++
	}

	// Get the elements from arr2 one by one.
	res := make([]int, 0, len(arr1))
	for _, i := range arr2 {
		count := counts[i]
		counts[i] = 0
		for j := 0; j < count; j++ {
			res = append(res, i)
		}
	}

	// Append the remaining numbers.
	for i, count := range counts {
		for j := 0; j < count; j++ {
			res = append(res, i)
		}
	}

	return res
}
