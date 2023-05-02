package chapter04

import "github.com/syhily/code-interviews/common"

func findIntersectionNode(f *common.ListNode, s *common.ListNode) *common.ListNode {
	fl, sl := lenOfListNode(f), lenOfListNode(s)

	if fl > sl {
		return innerFindTheIntersectionNode(f, s, fl-sl)
	} else {
		return innerFindTheIntersectionNode(s, f, sl-fl)
	}
}

func lenOfListNode(l *common.ListNode) int {
	var fl int
	for l != nil {
		fl++
		l = l.Next
	}
	return fl
}

func innerFindTheIntersectionNode(l *common.ListNode, s *common.ListNode, jump int) *common.ListNode {
	for i := 0; i < jump; i++ {
		l = l.Next
	}

	for l != nil && s != nil {
		if l == s {
			return l
		}

		l = l.Next
		s = s.Next
	}

	// Couldn't find the common node.
	return nil
}
