package chapter08

import "github.com/syhily/code-interviews/common"

func summarizeTheTreeNode(node *common.TreeNode[int]) *common.TreeNode[int] {
	stack, curr, sum := common.NewStack[*common.TreeNode[int]](), node, 0

	for curr != nil || !stack.Empty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Right
		}

		curr = stack.Pop()
		sum += curr.Value
		curr.Value = sum
		curr = curr.Left
	}

	return node
}
