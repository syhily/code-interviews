package chapter08

import "github.com/syhily/code-interviews/common"

func flattenTreeNode(node *common.TreeNode[int]) *common.TreeNode[int] {
	head, _ := innerFlattenTreeNode(node)
	return head
}

func innerFlattenTreeNode(node *common.TreeNode[int]) (*common.TreeNode[int], *common.TreeNode[int]) {
	if node == nil {
		return nil, nil
	}

	lh, lt := innerFlattenTreeNode(node.Left)
	if lh == nil {
		lh = node
		lt = node
	} else {
		lt.Right = node
		lt = node
	}
	node.Left = nil

	rh, rt := innerFlattenTreeNode(node.Right)
	if rh != nil {
		lt.Right = rh
	} else {
		rt = lt
	}

	return lh, rt
}
