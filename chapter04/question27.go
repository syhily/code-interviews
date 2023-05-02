package chapter04

import "github.com/syhily/code-interviews/common"

func isPalindromeListNode(l *common.ListNode) bool {
	if l == nil || l.Next == nil {
		return true
	}

	fast, slow := l, l
	var prev, right *common.ListNode

	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			break
		}
		prev = slow
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next
			right = slow.Next
		} else {
			right = slow
		}
	}

	prev.Next = nil
	left := reverseListNode(l)

	for left != nil {
		if left.Value != right.Value {
			return false
		}

		left = left.Next
		right = right.Next
	}

	return true
}
