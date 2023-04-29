package chapter07

import (
	"math"

	"github.com/syhily/code-interviews/common"
)

func maxValuesForTreeNode(tree *TreeNode) []int {
	if tree == nil {
		return nil
	}

	q := newQueue[*TreeNode]()
	q.add(tree)
	q.add(nil)

	var res []int
	curr := math.MinInt

	for !q.empty() {
		node := q.remove()

		if node == nil {
			res = append(res, curr)
			curr = math.MinInt
			if !q.empty() {
				q.add(nil)
			}
			continue
		}
		if node.left != nil {
			q.add(node.left)
		}
		if node.right != nil {
			q.add(node.right)
		}

		curr = common.MaxNumber(curr, node.value)
	}

	return res
}
