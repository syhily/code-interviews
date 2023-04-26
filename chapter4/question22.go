package chapter4

func findTheCircleInListNode(l *ListNode) *ListNode {
	// Guard logic, in case of the empty or single list node.
	if l == nil || l.next == nil {
		return nil
	}

	fast, slow := l, l

	for fast != nil {
		fast = fast.next

		// In case of there is no circle and the fast pointer hit the end.
		if fast != nil {
			fast = fast.next
		} else {
			return nil
		}

		slow = slow.next

		if slow == fast {
			break
		}
	}

	if fast == nil {
		// Didn't find the circle
		return nil
	}

	return slow
}
