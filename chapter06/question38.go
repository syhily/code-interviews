package chapter06

func dailyTemperatures(temperatures ...int) []int {
	s := newStack[int]()
	res := make([]int, len(temperatures))

	for i, temperature := range temperatures {
		if s.empty() {
			s.push(i)
			continue
		}

		for !s.empty() && temperatures[s.peek()] < temperature {
			idx := s.pop()
			res[idx] = i - idx
		}

		s.push(i)
	}

	for !s.empty() {
		res[s.pop()] = 0
	}

	return res
}
