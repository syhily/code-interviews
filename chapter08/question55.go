package chapter08

import "github.com/syhily/code-interviews/common"

type BSTIterator interface {
	next() int
	hasNext() bool
}

type bstIterator struct {
	stack common.Stack[*common.TreeNode[int]]
	head  *common.TreeNode[int]
}

func (b *bstIterator) next() int {
	for b.head != nil {
		b.stack.Push(b.head)
		b.head = b.head.Left
	}

	b.head = b.stack.Pop()
	res := b.head.Value
	b.head = b.head.Right

	return res
}

func (b *bstIterator) hasNext() bool {
	return b.head != nil || !b.stack.Empty()
}

func newBSTIterator(node *common.TreeNode[int]) BSTIterator {
	return &bstIterator{
		stack: common.NewStack[*common.TreeNode[int]](),
		head:  node,
	}
}
