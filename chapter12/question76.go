package chapter12

import (
	"crypto/rand"
	"math/big"
)

func quickSort(numbers ...int) []int {
	innerQuickSort(numbers, 0, len(numbers)-1)
	return numbers
}

func innerQuickSort(numbers []int, start, end int) {
	if start >= end {
		return
	}

	// Use a random value as the comparator.
	random, _ := rand.Int(rand.Reader, big.NewInt(int64(end-start+1)))
	pivot := int(random.Int64()) + start

	numbers[start], numbers[pivot] = numbers[pivot], numbers[start]
	pivot = start

	for i := start + 1; i <= end; i++ {
		if numbers[i] <= numbers[pivot] {
			numbers[pivot], numbers[i] = numbers[i], numbers[pivot]
			pivot++
			numbers[pivot], numbers[i] = numbers[i], numbers[pivot]
		}
	}

	innerQuickSort(numbers, start, pivot-1)
	innerQuickSort(numbers, pivot+1, end)
}
