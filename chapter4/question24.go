package chapter4

func reverseListNode(l *ListNode) *ListNode {
	c := l
	var prev *ListNode

	for c != nil {
		next := c.next
		c.next = prev
		prev = c
		c = next
	}

	return prev
}
