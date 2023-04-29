package chapter06

import "github.com/syhily/code-interviews/common"

func largestRectangleArea(heights []int) int {
	s := newStack[int]()
	s.push(-1)

	max := 0

	for i, h := range heights {
		for s.peek() != -1 && heights[s.peek()] >= h {
			height := heights[s.pop()]
			width := i - s.peek() - 1
			max = common.MaxNumber(max, width*height)
		}

		s.push(i)
	}

	for s.peek() != -1 {
		height := heights[s.pop()]
		width := len(heights) - s.peek() - 1
		max = common.MaxNumber(max, width*height)
	}

	return max
}
