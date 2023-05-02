package chapter04

import "github.com/syhily/code-interviews/common"

func reorderListNode(node *common.ListNode) *common.ListNode {
	if node == nil {
		return nil
	}

	fast, slow := node, node
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			break
		}

		slow = slow.Next
	}

	r, first, second := &common.ListNode{}, node, slow.Next
	slow.Next = nil
	second = reverseListNode(second)

	t := r
	for first != nil {
		t.Next = first
		t = t.Next
		first = first.Next

		if second != nil {
			t.Next = second
			t = t.Next
			second = second.Next
		}
	}

	return r.Next
}
