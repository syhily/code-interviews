package chapter08

import "github.com/syhily/code-interviews/common"

// Two sum implementation in BST
func twoSum(node *common.TreeNode[int], num int) bool {
	next := leftPointer(node)
	prev := rightPointer(node)

	if next.hasNext() && prev.hasPrev() {
		l, r := next.next(), prev.prev()

		for l < r {
			v := l + r
			if v < num {
				l = next.next()
			} else if v > num {
				r = prev.prev()
			} else {
				return true
			}
		}
	}

	return false
}

func leftPointer(node *common.TreeNode[int]) BSTIterator {
	return newBSTIterator(node)
}

func rightPointer(node *common.TreeNode[int]) ReservedBSTIterator {
	return &reservedBSTIterator{
		head:  node,
		stack: common.NewStack[*common.TreeNode[int]](),
	}
}

type ReservedBSTIterator interface {
	hasPrev() bool
	prev() int
}

type reservedBSTIterator struct {
	head  *common.TreeNode[int]
	stack common.Stack[*common.TreeNode[int]]
}

func (r *reservedBSTIterator) prev() int {
	for r.head != nil {
		r.stack.Push(r.head)
		r.head = r.head.Right
	}

	r.head = r.stack.Pop()
	val := r.head.Value
	r.head = r.head.Left

	return val
}

func (r *reservedBSTIterator) hasPrev() bool {
	return !r.stack.Empty() || r.head != nil
}
