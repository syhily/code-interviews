package chapter06

import "github.com/syhily/code-interviews/common"

func asteroidCollision(asteroids ...int) []int {
	s := common.NewStack[int]()

	for _, as := range asteroids {
		if s.Empty() || as >= 0 {
			s.Push(as)
			continue
		}

		for {
			if s.Empty() {
				s.Push(as)
				break
			}

			if head := s.Peek(); head >= 0 {
				if head > -as {
					break
				} else {
					s.Pop()
				}
			} else {
				s.Push(as)
				break
			}
		}
	}

	res := make([]int, s.Size())
	for i := s.Size() - 1; !s.Empty(); i-- {
		res[i] = s.Pop()
	}

	return res
}
