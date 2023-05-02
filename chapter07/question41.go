package chapter07

import "github.com/syhily/code-interviews/common"

type MovingAverage interface {
	next(val int) float64
}

type movingAverageSlice struct {
	elements []int
	start    int
	size     int
	sum      int
}

func (m *movingAverageSlice) next(val int) float64 {
	i := len(m.elements)
	if i < m.size {
		m.elements = append(m.elements, val)
	} else {
		m.sum -= m.elements[m.start]
		m.elements[m.start] = val
		m.start = (m.start + 1) % m.size
	}

	m.sum += val
	return float64(m.sum) / float64(len(m.elements))
}

func newMovingAverageSlice(size int) MovingAverage {
	return &movingAverageSlice{
		elements: make([]int, 0, size),
		start:    0,
		size:     size,
	}
}

type movingAverageQueue struct {
	queue common.Queue[int]
	sum   int
	size  int
}

func (m *movingAverageQueue) next(val int) float64 {
	if m.queue.Size() == m.size {
		m.sum -= m.queue.Remove()
	}

	m.queue.Add(val)
	m.sum += val

	return float64(m.sum) / float64(m.queue.Size())
}

func newMovingAverageQueue(size int) MovingAverage {
	return &movingAverageQueue{
		queue: common.NewQueue[int](),
		size:  size,
	}
}
