package common

type TreeNode[T comparable] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewTreeNode[T comparable](value T) *TreeNode[T] {
	return &TreeNode[T]{Value: value}
}

func InorderTraversal[T comparable](root *TreeNode[T], visit func(node *TreeNode[T])) {
	s, curr := NewStack[*TreeNode[T]](), root
	for !s.Empty() || curr != nil {
		for curr != nil {
			s.Push(curr)
			curr = curr.Left
		}

		t := s.Pop()
		visit(t)
		curr = t.Right
	}
}

func PreorderTraversal[T comparable](root *TreeNode[T], visit func(node *TreeNode[T])) {
	s, curr := NewStack[*TreeNode[T]](), root
	for !s.Empty() || curr != nil {
		for curr != nil {
			visit(curr)
			s.Push(curr)
			curr = curr.Left
		}

		curr = s.Pop()
		curr = curr.Right
	}
}

func PostorderTraversal[T comparable](root *TreeNode[T], visit func(node *TreeNode[T])) {
	s, curr := NewStack[*TreeNode[T]](), root
	var prev *TreeNode[T]
	for curr != nil || !s.Empty() {
		for curr != nil {
			s.Push(curr)
			curr = curr.Left
		}

		curr = s.Peek()
		if curr.Right != nil && curr.Right != prev {
			curr = curr.Right
		} else {
			s.Pop()
			visit(curr)
			prev = curr
			curr = nil
		}
	}
}
