package chapter04

func reorderListNode(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	fast, slow := node, node
	for fast != nil {
		fast = fast.next
		if fast != nil {
			fast = fast.next
		} else {
			break
		}

		slow = slow.next
	}

	r, first, second := &ListNode{}, node, slow.next
	slow.next = nil
	second = reverseListNode(second)

	t := r
	for first != nil {
		t.next = first
		t = t.next
		first = first.next

		if second != nil {
			t.next = second
			t = t.next
			second = second.next
		}
	}

	return r.next
}
