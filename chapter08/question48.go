package chapter08

import (
	"strconv"
	"strings"

	"github.com/syhily/code-interviews/common"
)

const (
	delimiter       = ","
	nilPlacerHolder = "#"
)

func serializeTreeNode(node *common.TreeNode[int]) string {
	return serialize(node, strconv.Itoa)
}

func deserializeTreeNode(text string) *common.TreeNode[int] {
	return deserialize(text, func(s string) int {
		r, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return r
	})
}

func serialize[T comparable](node *common.TreeNode[T], fn func(T) string) string {
	if node == nil {
		return nilPlacerHolder
	}

	return fn(node.Value) + delimiter + serialize(node.Left, fn) + delimiter + serialize(node.Right, fn)
}

func deserialize[T comparable](text string, fn func(string) T) *common.TreeNode[T] {
	nodes := strings.Split(text, delimiter)
	tree, _ := innerDeserialize(nodes, 0, fn)

	return tree
}

func innerDeserialize[T comparable](nodes []string, start int, fn func(string) T) (*common.TreeNode[T], int) {
	var node *common.TreeNode[T]

	if nodes[start] == nilPlacerHolder {
		return nil, start + 1
	} else {
		node = &common.TreeNode[T]{Value: fn(nodes[start])}
	}

	node.Left, start = innerDeserialize(nodes, start+1, fn)
	node.Right, start = innerDeserialize(nodes, start, fn)

	return node, start
}
