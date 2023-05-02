package chapter04

import "github.com/syhily/code-interviews/common"

func removeNthFromTheEnd(l *common.ListNode, n int) *common.ListNode {
	dummy := &common.ListNode{Next: l}

	first := l
	for i := 0; i < n; i++ {
		first = first.Next
	}

	second := dummy
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next

	return dummy.Next
}
