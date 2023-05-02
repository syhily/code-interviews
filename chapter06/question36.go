package chapter06

import (
	"strconv"

	"github.com/syhily/code-interviews/common"
)

var operations = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func evalRPN(tokens ...string) int {
	s := common.NewStack[int]()

	for _, t := range tokens {
		if operation, ok := operations[t]; ok {
			r := s.Pop()
			l := s.Pop()
			s.Push(operation(l, r))
		} else {
			i, _ := strconv.Atoi(t)
			s.Push(i)
		}
	}

	return s.Pop()
}
