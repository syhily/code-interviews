package chapter5

import "math/rand"

type RandomSet interface {
	insert(value int) bool
	remove(value int) bool
	getRandom() int
}

type randomSet struct {
	values   []int
	position map[int]int
}

func (r *randomSet) insert(value int) bool {
	if _, ok := r.position[value]; !ok {
		r.values = append(r.values, value)
		r.position[value] = len(r.values) - 1
		return true
	}

	return false
}

func (r *randomSet) remove(value int) bool {
	if d, ok := r.position[value]; ok {
		delete(r.position, value)
		l := len(r.values) - 1

		// Switch the value and delete the last value to get a O(1) operation.
		r.position[r.values[l]] = d
		r.values[d] = r.values[l]
		r.values = r.values[:l]

		return true
	}
	return false
}

func (r *randomSet) getRandom() int {
	i := rand.Intn(len(r.values))
	return r.values[i]
}
