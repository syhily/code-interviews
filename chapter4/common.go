package chapter4

// ListNode is used as the linked list. This struct is used in chapter 4.
type ListNode struct {
	value int
	next  *ListNode
}

func New(nodes ...int) *ListNode {
	l := &ListNode{value: nodes[0]}

	curr := l
	for _, node := range nodes[1:] {
		curr.next = &ListNode{value: node}
		curr = curr.next
	}

	return l
}

func NewCircleNode(start int, nodes ...int) (*ListNode, *ListNode) {
	var l, curr, s *ListNode

	for i, node := range nodes {
		if i == 0 {
			l = &ListNode{value: node}
			curr = l
			if node == start {
				s = l
			}
			continue
		}

		curr.next = &ListNode{value: node}
		curr = curr.next
		if start == node {
			s = curr
		}
	}

	if s != nil {
		curr.next = s
	}

	return l, s
}

func NewJoinNode(f, s, c []int) (*ListNode, *ListNode, *ListNode) {
	cc := New(c...)
	var fn, sn, cf, cs *ListNode

	for _, n := range f {
		if fn == nil {
			fn = &ListNode{value: n}
			cf = fn
			continue
		}

		cf.next = &ListNode{value: n}
		cf = cf.next
	}

	for _, n := range s {
		if sn == nil {
			sn = &ListNode{value: n}
			cs = sn
			continue
		}

		cs.next = &ListNode{value: n}
		cs = cs.next
	}

	cf.next = cc
	cs.next = cc

	return fn, sn, cc
}

func appendNode(l *ListNode, value int) *ListNode {
	if l == nil {
		return &ListNode{value: value}
	}

	c := l
	for c.next != nil {
		c = c.next
	}
	c.next = &ListNode{value: value}

	return l
}

func dummyAppendNode(l *ListNode, value int) *ListNode {
	dummy := &ListNode{next: l}

	tail := dummy
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = &ListNode{value: value}

	return dummy.next
}

func deleteNode(l *ListNode, value int) *ListNode {
	if l == nil {
		return l
	}

	if l.value == value {
		return l.next
	}

	curr := l
	for curr.next != nil {
		if curr.next.value == value {
			curr.next = curr.next.next
			break
		}
	}

	return l
}

func dummyDeleteNode(l *ListNode, value int) *ListNode {
	dummy := &ListNode{next: l}

	curr := dummy
	for curr.next != nil {
		if curr.next.value == value {
			curr.next = curr.next.next
			break
		}
	}

	return dummy.next
}
