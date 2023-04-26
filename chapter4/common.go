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
