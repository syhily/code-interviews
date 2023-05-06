package chapter08

import (
	"math"

	"github.com/syhily/code-interviews/common"
)

func findTheMaxRoutine(node *common.TreeNode[int]) int {
	max := math.MinInt
	innerFindTheMaxRoutine(node, &max)

	return max
}

func innerFindTheMaxRoutine(node *common.TreeNode[int], max *int) int {
	if node == nil {
		return 0
	}

	left := innerFindTheMaxRoutine(node.Left, max)
	right := innerFindTheMaxRoutine(node.Right, max)

	*max = common.MaxNumber(*max, left)
	*max = common.MaxNumber(*max, right)
	*max = common.MaxNumber(*max, left+right+node.Value)

	return node.Value + common.MaxNumber(left, right)
}
