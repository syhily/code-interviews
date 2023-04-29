package chapter07

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
	queue Queue[int]
	sum   int
	size  int
}

func (m *movingAverageQueue) next(val int) float64 {
	if m.queue.size() == m.size {
		m.sum -= m.queue.remove()
	}

	m.queue.add(val)
	m.sum += val

	return float64(m.sum) / float64(m.queue.size())
}

func newMovingAverageQueue(size int) MovingAverage {
	return &movingAverageQueue{
		queue: newQueue[int](),
		size:  size,
	}
}
