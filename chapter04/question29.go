package chapter04

func insertNodeIntoCircle(node *ListNode, value int) *ListNode {
	if node == nil {
		node = &ListNode{value: value}
		node.next = node
		return node
	}

	start, max := node.next, node
	n := &ListNode{value: value}

	for start != node {
		if start.value > max.value {
			max = start
		}

		if value > start.value && value < start.next.value {
			insertAfterNode(start, n)
			return node
		}

		start = start.next
	}

	// Couldn't find the right position, insert after the max node.
	insertAfterNode(max, n)

	return node
}

func insertAfterNode(start, node *ListNode) {
	next := start.next
	start.next = node
	node.next = next
}
