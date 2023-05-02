package chapter12

import "github.com/syhily/code-interviews/common"

func mergeKListNodes(nodes ...*common.ListNode) *common.ListNode {
	length := len(nodes)

	if length == 0 {
		return nil
	} else if length == 1 {
		return nodes[0]
	}

	middle := length / 2

	left := mergeKListNodes(nodes[:middle]...)
	right := mergeKListNodes(nodes[middle:]...)

	return mergeList(left, right)
}
