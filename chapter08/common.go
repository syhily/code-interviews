package chapter08

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func newNode(value int) *TreeNode {
	return &TreeNode{value: value}
}
