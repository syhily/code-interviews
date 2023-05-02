package common

import "fmt"

var ErrNoElementInQueue = fmt.Errorf("there is no elements in the queue")

// Queue is a common interface for representing the FIFO data structure.
type Queue[V comparable] interface {
	Add(v V)
	Remove() V
	Element() V
	Empty() bool
	Size() int
}

type element[V comparable] struct {
	value V
	next  *element[V]
}

type queue[V comparable] struct {
	head   *element[V]
	tail   *element[V]
	counts int
}

func (q *queue[V]) Add(v V) {
	e := &element[V]{value: v}
	if q.head == nil {
		q.head = e
		q.tail = q.head
	} else {
		q.tail.next = e
		q.tail = e
	}
	q.counts++
}

func (q *queue[V]) Remove() V {
	head := q.head
	if head == nil {
		panic(ErrNoElementInQueue)
	}
	q.head = head.next
	q.counts--
	return head.value
}

func (q *queue[V]) Element() V {
	head := q.head
	if head == nil {
		panic(ErrNoElementInQueue)
	}
	return head.value
}

func (q *queue[V]) Empty() bool {
	return q.counts == 0
}

func (q *queue[V]) Size() int {
	return q.counts
}

func NewQueue[V comparable]() Queue[V] {
	return &queue[V]{}
}
