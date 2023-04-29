package chapter07

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
	queue Queue[*TreeNode]
}

func (c *cbtInserter) insert(value int) *TreeNode {
	parent := c.queue.element()
	node := &TreeNode{value: value}

	if parent.left == nil {
		parent.left = node
		c.queue.add(node)
	} else {
		parent.right = node
		c.queue.add(node)
		c.queue.remove()
	}

	return parent
}

func (c *cbtInserter) root() *TreeNode {
	return c.tree
}

func newCBTInserter(node *TreeNode) CBTInserter {
	q := newQueue[*TreeNode]()
	q.add(node)

	for {
		n := q.element()
		if n.left == nil || n.right == nil {
			break
		}

		q.add(n.left)
		q.add(n.right)
		q.remove()
	}

	return &cbtInserter{
		tree:  node,
		queue: q,
	}
}
