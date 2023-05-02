package chapter04

import "github.com/syhily/code-interviews/common"

func addTwoNumbers(node1, node2 *common.ListNode) *common.ListNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	e, s1, s2 := 0, reverseListNode(node1), reverseListNode(node2)
	var r, s3 *common.ListNode

	divideNumber := func(num int) (int, int) {
		return num % 10, num / 10
	}

	for s1 != nil || s2 != nil {
		v1, v2, v3 := 0, 0, 0
		if s1 != nil {
			v1 = s1.Value
			s1 = s1.Next
		}
		if s2 != nil {
			v2 = s2.Value
			s2 = s2.Next
		}

		v3, e = divideNumber(v1 + v2 + e)

		if r == nil {
			r = &common.ListNode{Value: v3}
			s3 = r
		} else {
			s3.Next = &common.ListNode{Value: v3}
			s3 = s3.Next
		}
	}

	return reverseListNode(r)
}
