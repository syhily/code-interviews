package chapter06

import (
	"fmt"
)

var ErrNoElementInStack = fmt.Errorf("there is no elements in the stack")

type Entry[V comparable] struct {
	value V
	next  *Entry[V]
}

// Stack is a common interface for representing the FILO data structure.
type Stack[V comparable] interface {
	push(v V)
	pop() V
	peek() V
	empty() bool
	size() int
}

type stack[V comparable] struct {
	entry  *Entry[V]
	counts int
}

func (s *stack[V]) empty() bool {
	return s.entry == nil
}

func (s *stack[V]) size() int {
	return s.counts
}

func (s *stack[V]) push(v V) {
	if s.entry == nil {
		s.entry = &Entry[V]{value: v}
	} else {
		e := &Entry[V]{value: v}
		e.next = s.entry
		s.entry = e
	}
	s.counts++
}

func (s *stack[V]) pop() V {
	if s.entry == nil {
		panic(ErrNoElementInStack)
	}
	e := s.entry
	s.entry = e.next
	s.counts--

	return e.value
}

func (s *stack[V]) peek() V {
	if s.entry == nil {
		panic(ErrNoElementInStack)
	}

	return s.entry.value
}

func newStack[V comparable]() Stack[V] {
	return &stack[V]{}
}
