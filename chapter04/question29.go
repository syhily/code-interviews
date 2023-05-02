package chapter04

import "github.com/syhily/code-interviews/common"

func insertNodeIntoCircle(node *common.ListNode, value int) *common.ListNode {
	if node == nil {
		node = &common.ListNode{Value: value}
		node.Next = node
		return node
	}

	start, max := node.Next, node
	n := &common.ListNode{Value: value}

	for start != node {
		if start.Value > max.Value {
			max = start
		}

		if value > start.Value && value < start.Next.Value {
			insertAfterNode(start, n)
			return node
		}

		start = start.Next
	}

	// Couldn't find the right position, insert after the max node.
	insertAfterNode(max, n)

	return node
}

func insertAfterNode(start, node *common.ListNode) {
	next := start.Next
	start.Next = node
	node.Next = next
}
