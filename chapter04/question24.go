package chapter04

import "github.com/syhily/code-interviews/common"

func reverseListNode(l *common.ListNode) *common.ListNode {
	c := l
	var prev *common.ListNode

	for c != nil {
		next := c.Next
		c.Next = prev
		prev = c
		c = next
	}

	return prev
}
