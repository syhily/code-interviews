package chapter06

import "github.com/syhily/code-interviews/common"

func largestRectangleArea(heights []int) int {
	s := common.NewStack[int]()
	s.Push(-1)

	max := 0

	for i, h := range heights {
		for s.Peek() != -1 && heights[s.Peek()] >= h {
			height := heights[s.Pop()]
			width := i - s.Peek() - 1
			max = common.MaxNumber(max, width*height)
		}

		s.Push(i)
	}

	for s.Peek() != -1 {
		height := heights[s.Pop()]
		width := len(heights) - s.Peek() - 1
		max = common.MaxNumber(max, width*height)
	}

	return max
}
