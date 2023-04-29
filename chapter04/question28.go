package chapter04

type Node struct {
	value int
	prev  *Node
	next  *Node
	child *Node
}

func flattenNode(node *Node) *Node {
	n, _ := innerFlattenNode(node)
	return n
}

func innerFlattenNode(node *Node) (*Node, *Node) {
	iter, start := node, &Node{}
	end := start

	for iter != nil {
		end.next = iter
		iter.prev = end
		end = end.next

		if iter.child != nil {
			s, e := innerFlattenNode(iter.child)
			iter.child = nil
			next := iter.next
			end.next = s
			s.prev = end
			end = e
			iter = next
			continue
		}

		iter = iter.next
	}

	next := start.next
	next.prev = nil

	return next, end
}

type CreateNode struct {
	value int
	prev  int
	next  int
	child int
}

func createNode(nodes ...CreateNode) *Node {
	var start *Node
	idx := make(map[int]*Node)

	for _, n := range nodes {
		node := &Node{value: n.value}
		idx[n.value] = node
		if start == nil {
			start = node
		}
	}

	for _, n := range nodes {
		node := idx[n.value]
		if n.next > 0 {
			next := idx[n.next]
			node.next = next
		}
		if n.prev > 0 {
			prev := idx[n.prev]
			node.prev = prev
		}
		if n.child > 0 {
			child := idx[n.child]
			node.child = child
		}
	}

	return start
}
