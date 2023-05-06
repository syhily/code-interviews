package chapter08

import (
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_newBSTIterator(t *testing.T) {
	type args struct {
		node *common.TreeNode[int]
	}
	type want struct {
		value int
	}
	tests := []struct {
		name  string
		args  args
		wants []want
	}{
		{
			name: "bst iterator",
			args: args{node: &common.TreeNode[int]{
				Value: 10,
				Left: &common.TreeNode[int]{
					Value: 8,
					Left: &common.TreeNode[int]{
						Value: 7,
					},
					Right: &common.TreeNode[int]{
						Value: 9,
					},
				},
				Right: &common.TreeNode[int]{
					Value: 12,
					Left: &common.TreeNode[int]{
						Value: 11,
					},
					Right: &common.TreeNode[int]{
						Value: 13,
					},
				},
			}},
			wants: []want{
				{value: 7},
				{value: 8},
				{value: 9},
				{value: 10},
				{value: 11},
				{value: 12},
				{value: 13},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := newBSTIterator(tt.args.node)
			for _, w := range tt.wants {
				if !iterator.hasNext() {
					t.Errorf("iterator.hasNext() should return true")
				}
				next := iterator.next()
				if next != w.value {
					t.Errorf("iterator.next() = %v, want %v", next, w.value)
				}
			}

			if iterator.hasNext() {
				t.Errorf("iterator.hasNext() should return false")
			}
		})
	}
}
