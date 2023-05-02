package chapter06

import "github.com/syhily/code-interviews/common"

func dailyTemperatures(temperatures ...int) []int {
	s := common.NewStack[int]()
	res := make([]int, len(temperatures))

	for i, temperature := range temperatures {
		if s.Empty() {
			s.Push(i)
			continue
		}

		for !s.Empty() && temperatures[s.Peek()] < temperature {
			idx := s.Pop()
			res[idx] = i - idx
		}

		s.Push(i)
	}

	for !s.Empty() {
		res[s.Pop()] = 0
	}

	return res
}
