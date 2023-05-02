package common

// ListNode is used as the linked list. This struct is used in chapter 4.
type ListNode struct {
	Value int
	Next  *ListNode
}

func NewListNode(nodes ...int) *ListNode {
	l := &ListNode{Value: nodes[0]}

	curr := l
	for _, node := range nodes[1:] {
		curr.Next = &ListNode{Value: node}
		curr = curr.Next
	}

	return l
}

func NewCircleNode(start int, nodes ...int) (l *ListNode, s *ListNode) {
	var curr *ListNode

	for i, node := range nodes {
		if i == 0 {
			l = &ListNode{Value: node}
			curr = l
			if node == start {
				s = l
			}
			continue
		}

		curr.Next = &ListNode{Value: node}
		curr = curr.Next
		if start == node {
			s = curr
		}
	}

	if s != nil {
		curr.Next = s
	}

	return
}

func NewJoinNode(f, s, c []int) (fn *ListNode, sn *ListNode, cc *ListNode) {
	cc = NewListNode(c...)
	var cf, cs *ListNode

	for _, n := range f {
		if fn == nil {
			fn = &ListNode{Value: n}
			cf = fn
			continue
		}

		cf.Next = &ListNode{Value: n}
		cf = cf.Next
	}

	for _, n := range s {
		if sn == nil {
			sn = &ListNode{Value: n}
			cs = sn
			continue
		}

		cs.Next = &ListNode{Value: n}
		cs = cs.Next
	}

	cf.Next = cc
	cs.Next = cc

	return
}

func appendNode(l *ListNode, value int) *ListNode {
	if l == nil {
		return &ListNode{Value: value}
	}

	c := l
	for c.Next != nil {
		c = c.Next
	}
	c.Next = &ListNode{Value: value}

	return l
}

func dummyAppendNode(l *ListNode, value int) *ListNode {
	dummy := &ListNode{Next: l}

	tail := dummy
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = &ListNode{Value: value}

	return dummy.Next
}

func deleteNode(l *ListNode, value int) *ListNode {
	if l == nil {
		return nil
	}

	if l.Value == value {
		return l.Next
	}

	curr := l
	for curr.Next != nil {
		if curr.Next.Value == value {
			curr.Next = curr.Next.Next
			break
		}
	}

	return l
}

func dummyDeleteNode(l *ListNode, value int) *ListNode {
	dummy := &ListNode{Next: l}

	curr := dummy
	for curr.Next != nil {
		if curr.Next.Value == value {
			curr.Next = curr.Next.Next
			break
		}
	}

	return dummy.Next
}
