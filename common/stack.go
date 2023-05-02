package common

import (
	"fmt"
)

var ErrNoElementInStack = fmt.Errorf("there is no elements in the stack")

// Stack is a common interface for representing the FILO data structure.
type Stack[V comparable] interface {
	Push(v V)
	Pop() V
	Peek() V
	Empty() bool
	Size() int
}

type entry[V comparable] struct {
	value V
	next  *entry[V]
}

type stack[V comparable] struct {
	entry  *entry[V]
	counts int
}

func (s *stack[V]) Empty() bool {
	return s.entry == nil
}

func (s *stack[V]) Size() int {
	return s.counts
}

func (s *stack[V]) Push(v V) {
	if s.entry == nil {
		s.entry = &entry[V]{value: v}
	} else {
		e := &entry[V]{value: v}
		e.next = s.entry
		s.entry = e
	}
	s.counts++
}

func (s *stack[V]) Pop() V {
	if s.entry == nil {
		panic(ErrNoElementInStack)
	}
	e := s.entry
	s.entry = e.next
	s.counts--

	return e.value
}

func (s *stack[V]) Peek() V {
	if s.entry == nil {
		panic(ErrNoElementInStack)
	}

	return s.entry.value
}

func NewStack[V comparable]() Stack[V] {
	return &stack[V]{}
}
