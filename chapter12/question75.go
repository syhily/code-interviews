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
	// Find the min and max value from arr1.
	max, min := math.MinInt, math.MaxInt
	for _, i := range arr1 {
		max = common.MaxNumber(max, i)
		min = common.MinNumber(min, i)
	}

	// Count the occur times of items in arr1.
	counts := make([]int, max-min+1)
	for _, i := range arr1 {
		counts[i-min]++
	}

	// Get the elements from arr2 one by one.
	res := make([]int, 0, len(arr1))
	for _, i := range arr2 {
		count := counts[i-min]
		counts[i-min] = 0
		for j := 0; j < count; j++ {
			res = append(res, i)
		}
	}

	// Append the remaining numbers.
	for i, count := range counts {
		for j := 0; j < count; j++ {
			res = append(res, i+min)
		}
	}

	return res
}
