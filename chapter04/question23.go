package chapter04

func findIntersectionNode(f *ListNode, s *ListNode) *ListNode {
	fl, sl := lenOfListNode(f), lenOfListNode(s)

	if fl > sl {
		return innerFindTheIntersectionNode(f, s, fl-sl)
	} else {
		return innerFindTheIntersectionNode(s, f, sl-fl)
	}
}

func lenOfListNode(l *ListNode) int {
	var fl int
	for l != nil {
		fl++
		l = l.next
	}
	return fl
}

func innerFindTheIntersectionNode(l *ListNode, s *ListNode, jump int) *ListNode {
	for i := 0; i < jump; i++ {
		l = l.next
	}

	for l != nil && s != nil {
		if l == s {
			return l
		}

		l = l.next
		s = s.next
	}

	// Couldn't find the common node.
	return nil
}
