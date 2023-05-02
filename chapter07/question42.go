package chapter07

import "github.com/syhily/code-interviews/common"

type RecentAverage interface {
	ping(time int) int
}

type recentAverage struct {
	queue     common.Queue[int]
	timeRange int
}

func newRecentAverage(r int) RecentAverage {
	return &recentAverage{
		queue:     common.NewQueue[int](),
		timeRange: r,
	}
}

func (r *recentAverage) ping(time int) int {
	if r.queue.Empty() {
		r.queue.Add(time)
		return 1
	}

	for !r.queue.Empty() && r.queue.Element()+r.timeRange < time {
		r.queue.Remove()
	}
	r.queue.Add(time)

	return r.queue.Size()
}
