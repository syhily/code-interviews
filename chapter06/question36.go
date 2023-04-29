package chapter06

import (
	"fmt"
	"strconv"
)

var operations = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func invalidRPN(err error, tokens ...string) {
	if err != nil {
		panic(fmt.Sprintf("invalid rpn expression: %v", tokens))
	}
}

func evalRPN(tokens ...string) int {
	s := newStack[int]()

	for _, t := range tokens {
		if operation, ok := operations[t]; ok {
			r, err := s.pop()
			invalidRPN(err, tokens...)

			l, err := s.pop()
			invalidRPN(err, tokens...)

			r = operation(l, r)
			s.push(r)
		} else {
			i, _ := strconv.Atoi(t)
			s.push(i)
		}
	}

	r, err := s.pop()
	invalidRPN(err, tokens...)

	return r
}
