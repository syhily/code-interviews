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

var (
	atoi = func(s string) int {
		r, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return r
	}
)

func serializeTreeNode(node *common.TreeNode[int]) string {
	return serialize(node, strconv.Itoa)
}

func deserializeTreeNode(text string) *common.TreeNode[int] {
	return deserialize(text, atoi)
}

func serialize[T comparable](node *common.TreeNode[T], fn func(T) string) string {
	if node == nil {
		return nilPlacerHolder
	}

	return fn(node.Value) + delimiter + serialize(node.Left, fn) + delimiter + serialize(node.Right, fn)
}

func deserialize[T comparable](text string, fn func(string) T) *common.TreeNode[T] {
	nodes, start := strings.Split(text, delimiter), 0
	return innerDeserialize(nodes, &start, fn)
}

func innerDeserialize[T comparable](nodes []string, start *int, fn func(string) T) *common.TreeNode[T] {
	var node *common.TreeNode[T]
	text := nodes[*start]
	*start++

	if text == nilPlacerHolder {
		return nil
	} else {
		node = &common.TreeNode[T]{Value: fn(text)}
	}

	node.Left = innerDeserialize(nodes, start, fn)
	node.Right = innerDeserialize(nodes, start, fn)

	return node
}
