package chapter08

import "github.com/syhily/code-interviews/common"

func countTheRoutine(node *common.TreeNode[int], desired int) int {
	return innerCountTheRoutine(node, desired, 0, map[int]int{0: 1})
}

func innerCountTheRoutine(node *common.TreeNode[int], desired, path int, paths map[int]int) int {
	if node == nil {
		return 0
	}

	path += node.Value
	count := paths[path-desired]

	paths[path]++
	count += innerCountTheRoutine(node.Left, desired, path, paths)
	count += innerCountTheRoutine(node.Right, desired, path, paths)
	paths[path]--

	return count
}
