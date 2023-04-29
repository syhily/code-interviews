package chapter07

import "fmt"

var ErrNoElementInQueue = fmt.Errorf("there is no elements in the queue")

// Queue is a common interface for representing the FIFO data structure.
type Queue[V comparable] interface {
	add(v V)
	remove() V
	element() V
	empty() bool
	size() int
}

type entry[V comparable] struct {
	value V
	next  *entry[V]
}

type queue[V comparable] struct {
	head   *entry[V]
	tail   *entry[V]
	counts int
}

func (q *queue[V]) add(v V) {
	e := &entry[V]{value: v}
	if q.head == nil {
		q.head = e
		q.tail = q.head
	} else {
		q.tail.next = e
		q.tail = e
	}
	q.counts++
}

func (q *queue[V]) remove() V {
	head := q.head
	if head == nil {
		panic(ErrNoElementInQueue)
	}
	q.head = head.next
	q.counts--
	return head.value
}

func (q *queue[V]) element() V {
	head := q.head
	if head == nil {
		panic(ErrNoElementInQueue)
	}
	return head.value
}

func (q *queue[V]) empty() bool {
	return q.counts == 0
}

func (q *queue[V]) size() int {
	return q.counts
}

func newQueue[V comparable]() Queue[V] {
	return &queue[V]{}
}
