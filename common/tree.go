package common

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{value: value}
}
