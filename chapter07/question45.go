package chapter07

import (
	"fmt"
	"math"
)

func leftValueFromLowestLayer(tree *TreeNode) int {
	if tree == nil {
		panic(fmt.Errorf("tree parameter shouldn't be nil"))
	}

	q := newQueue[*TreeNode]()
	q.add(tree)
	q.add(nil)

	start, left := true, math.MinInt

	for !q.empty() {
		node := q.remove()

		if node == nil {
			if !q.empty() {
				q.add(nil)
				start = true
			}
			continue
		}

		if node.left != nil {
			q.add(node.left)
		}
		if node.right != nil {
			q.add(node.right)
		}
		if start {
			left = node.value
			start = false
		}
	}

	return left
}
