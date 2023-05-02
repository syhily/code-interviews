package chapter07

import "github.com/syhily/code-interviews/common"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type CBTInserter interface {
	insert(value int) *TreeNode
	root() *TreeNode
}

type cbtInserter struct {
	tree  *TreeNode
	queue common.Queue[*TreeNode]
}

func (c *cbtInserter) insert(value int) *TreeNode {
	parent := c.queue.Element()
	node := &TreeNode{value: value}

	if parent.left == nil {
		parent.left = node
		c.queue.Add(node)
	} else {
		parent.right = node
		c.queue.Add(node)
		c.queue.Remove()
	}

	return parent
}

func (c *cbtInserter) root() *TreeNode {
	return c.tree
}

func newCBTInserter(node *TreeNode) CBTInserter {
	q := common.NewQueue[*TreeNode]()
	q.Add(node)

	for {
		n := q.Element()
		if n.left == nil || n.right == nil {
			break
		}

		q.Add(n.left)
		q.Add(n.right)
		q.Remove()
	}

	return &cbtInserter{
		tree:  node,
		queue: q,
	}
}
