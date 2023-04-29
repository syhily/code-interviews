package chapter06

func asteroidCollision(asteroids ...int) []int {
	s := newStack[int]()

	for _, as := range asteroids {
		if s.empty() || as >= 0 {
			s.push(as)
			continue
		}

		for {
			if s.empty() {
				s.push(as)
				break
			}

			if head := s.peek(); head >= 0 {
				if head > -as {
					break
				} else {
					s.pop()
				}
			} else {
				s.push(as)
				break
			}
		}
	}

	res := make([]int, s.size())
	for i := s.size() - 1; !s.empty(); i-- {
		res[i] = s.pop()
	}

	return res
}
