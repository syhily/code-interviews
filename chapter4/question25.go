package chapter4

func addTwoNumbers(node1, node2 *ListNode) *ListNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	e, s1, s2 := 0, reverseListNode(node1), reverseListNode(node2)
	var r, s3 *ListNode

	divideNumber := func(num int) (int, int) {
		return num % 10, num / 10
	}

	for s1 != nil || s2 != nil {
		v1, v2, v3 := 0, 0, 0
		if s1 != nil {
			v1 = s1.value
			s1 = s1.next
		}
		if s2 != nil {
			v2 = s2.value
			s2 = s2.next
		}

		v3, e = divideNumber(v1 + v2 + e)

		if r == nil {
			r = &ListNode{value: v3}
			s3 = r
		} else {
			s3.next = &ListNode{value: v3}
			s3 = s3.next
		}
	}

	return reverseListNode(r)
}
