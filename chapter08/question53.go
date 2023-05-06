package chapter08

import "github.com/syhily/code-interviews/common"

func findTheNextNode(node *common.TreeNode[int], number int) *common.TreeNode[int] {
	curr := node
	var prev *common.TreeNode[int]

	for curr != nil {
		if curr.Value <= number {
			curr = curr.Right
		} else {
			prev = curr
			curr = curr.Left
		}
	}

	return prev
}
