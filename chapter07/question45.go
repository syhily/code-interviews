package chapter07

import (
	"fmt"
	"math"

	"github.com/syhily/code-interviews/common"
)

func leftValueFromLowestLayer(tree *TreeNode) int {
	if tree == nil {
		panic(fmt.Errorf("tree parameter shouldn't be nil"))
	}

	q := common.NewQueue[*TreeNode]()
	q.Add(tree)
	q.Add(nil)

	start, left := true, math.MinInt

	for !q.Empty() {
		node := q.Remove()

		if node == nil {
			if !q.Empty() {
				q.Add(nil)
				start = true
			}
			continue
		}

		if node.left != nil {
			q.Add(node.left)
		}
		if node.right != nil {
			q.Add(node.right)
		}
		if start {
			left = node.value
			start = false
		}
	}

	return left
}
