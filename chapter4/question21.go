package chapter4

func removeNthFromTheEnd(l *ListNode, n int) *ListNode {
	dummy := &ListNode{next: l}

	first := l
	for i := 0; i < n; i++ {
		first = first.next
	}

	second := dummy
	for first != nil {
		first = first.next
		second = second.next
	}
	second.next = second.next.next

	return dummy.next
}
