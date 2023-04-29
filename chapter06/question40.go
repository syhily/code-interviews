package chapter06

import "github.com/syhily/code-interviews/common"

func maximumRectangle(matrix [][]bool) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	heights := make([]int, len(matrix[0]))
	rectangle := 0

	for _, row := range matrix {
		for i := 0; i < len(heights); i++ {
			if row[i] {
				heights[i]++
			} else {
				heights[i] = 0
			}
		}

		rectangle = common.MaxNumber(rectangle, largestRectangleArea(heights))
	}

	return rectangle
}
