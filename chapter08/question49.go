package chapter08

import "github.com/syhily/code-interviews/common"

func calculateNumbers(nodes *common.TreeNode[int]) int {
	return innerCalculateNumbers(nodes, 0)
}

func innerCalculateNumbers(node *common.TreeNode[int], val int) int {
	if node == nil {
		return 0
	}

	val = val*10 + node.Value
	if node.Left == nil && node.Right == nil {
		return val
	}

	return innerCalculateNumbers(node.Left, val) + innerCalculateNumbers(node.Right, val)
}
