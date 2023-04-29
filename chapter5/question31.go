package chapter5

type LRUCache[K comparable, V any] interface {
	get(key K) V
	put(key K, value V)
}

type entry[K comparable, V any] struct {
	prev  *entry[K, V]
	next  *entry[K, V]
	key   K
	value V
}

type list[K comparable, V any] struct {
	head *entry[K, V]
	end  *entry[K, V]
}

func (l *list[K, V]) delete(e *entry[K, V]) {
	prev := e.prev
	next := e.next

	if prev == nil {
		next.prev = nil
		l.head = next
		return
	}

	if next == nil {
		prev.next = nil
		l.end = prev
		return
	}

	prev.next = next
	next.prev = prev
}

func (l *list[K, V]) append(k K, v V) *entry[K, V] {
	n := &entry[K, V]{key: k, value: v}

	if l.head == nil {
		l.head = n
		l.end = n
		return n
	}

	l.end.next = n
	n.prev = l.end
	l.end = n

	return n
}

func (l *list[K, V]) renew(e *entry[K, V]) {
	l.delete(e)
	l.end.next = e
	e.next = nil
	e.prev = l.end
	l.end = e
}

func (l *list[K, V]) expire() *entry[K, V] {
	head := l.head
	next := head.next

	if next != nil {
		l.head = next
		next.prev = nil
	} else {
		l.head = nil
		l.end = nil
	}

	return head
}

func newLRUCache[K comparable, V any](capacity int, defaults V) LRUCache[K, V] {
	return &lruCache[K, V]{
		capacity: capacity,
		defaults: defaults,
		nodes:    map[K]*entry[K, V]{},
		list:     &list[K, V]{},
	}
}

type lruCache[K comparable, V any] struct {
	capacity int
	defaults V
	nodes    map[K]*entry[K, V]
	list     *list[K, V]
}

func (l *lruCache[K, V]) get(k K) V {
	if e, ok := l.nodes[k]; ok {
		return e.value
	} else {
		return l.defaults
	}
}

func (l *lruCache[K, V]) put(k K, v V) {
	// Remove the oldest key if it exceeds the cache capacity.
	if l.capacity <= len(l.nodes) {
		e := l.list.expire()
		delete(l.nodes, e.key)
	}

	// Insert or update the existing entry.
	if e, ok := l.nodes[k]; ok {
		e.value = v
	} else {
		e = l.list.append(k, v)
		l.nodes[k] = e
	}
}
