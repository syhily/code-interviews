package chapter04

func isPalindromeListNode(l *ListNode) bool {
	if l == nil || l.next == nil {
		return true
	}

	fast, slow := l, l
	var prev, right *ListNode

	for fast != nil {
		if fast.next != nil {
			fast = fast.next
		} else {
			break
		}
		prev = slow
		slow = slow.next
		if fast.next != nil {
			fast = fast.next
			right = slow.next
		} else {
			right = slow
		}
	}

	prev.next = nil
	left := reverseListNode(l)

	for left != nil {
		if left.value != right.value {
			return false
		}

		left = left.next
		right = right.next
	}

	return true
}
