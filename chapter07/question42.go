package chapter07

type RecentAverage interface {
	ping(time int) int
}

type recentAverage struct {
	queue     Queue[int]
	timeRange int
}

func newRecentAverage(r int) RecentAverage {
	return &recentAverage{
		queue:     newQueue[int](),
		timeRange: r,
	}
}

func (r *recentAverage) ping(time int) int {
	if r.queue.empty() {
		r.queue.add(time)
		return 1
	}

	for !r.queue.empty() && r.queue.element()+r.timeRange < time {
		r.queue.remove()
	}
	r.queue.add(time)

	return r.queue.size()
}
