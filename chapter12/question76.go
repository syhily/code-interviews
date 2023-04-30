package chapter12

import (
	"crypto/rand"
	"fmt"
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

func findTheKth(k int, numbers ...int) int {
	l := len(numbers)
	if k > l {
		panic(fmt.Errorf("invalid k, it exceed the array length  %v", l))
	}

	return innerFindTheKth(numbers, 0, l-1, l-k)
}

func innerFindTheKth(numbers []int, start, end, des int) int {
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

	if pivot > des {
		return innerFindTheKth(numbers, start, pivot-1, des)
	} else if pivot < des {
		return innerFindTheKth(numbers, pivot+1, end, des)
	} else {
		return numbers[pivot]
	}
}
