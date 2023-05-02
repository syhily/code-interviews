package chapter07

import (
	"math"

	"github.com/syhily/code-interviews/common"
)

func maxValuesForTreeNode(tree *TreeNode) []int {
	if tree == nil {
		return nil
	}

	q := common.NewQueue[*TreeNode]()
	q.Add(tree)
	q.Add(nil)

	var res []int
	curr := math.MinInt

	for !q.Empty() {
		node := q.Remove()

		if node == nil {
			res = append(res, curr)
			curr = math.MinInt
			if !q.Empty() {
				q.Add(nil)
			}
			continue
		}
		if node.left != nil {
			q.Add(node.left)
		}
		if node.right != nil {
			q.Add(node.right)
		}

		curr = common.MaxNumber(curr, node.value)
	}

	return res
}
