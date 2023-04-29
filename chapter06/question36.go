package chapter06

import (
	"strconv"
)

var operations = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func evalRPN(tokens ...string) int {
	s := newStack[int]()

	for _, t := range tokens {
		if operation, ok := operations[t]; ok {
			r := s.pop()
			l := s.pop()
			s.push(operation(l, r))
		} else {
			i, _ := strconv.Atoi(t)
			s.push(i)
		}
	}

	return s.pop()
}
