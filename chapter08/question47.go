package chapter08

import "github.com/syhily/code-interviews/common"

func removeZeroNode(node *common.TreeNode[int]) *common.TreeNode[int] {
	if node == nil {
		return nil
	}

	node.Left = removeZeroNode(node.Left)
	node.Right = removeZeroNode(node.Right)

	if node.Left == nil && node.Right == nil && node.Value == 0 {
		return nil
	} else {
		return node
	}
}
