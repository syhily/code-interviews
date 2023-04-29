package chapter07

import (
	"reflect"
	"testing"
)

func Test_rightViewOfTreeNode(t *testing.T) {
	type args struct {
		tree *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "calculate right view of the tree 1",
			args: args{tree: &TreeNode{
				value: 8,
				left: &TreeNode{
					value: 6,
				},
				right: &TreeNode{
					value: 10,
					left: &TreeNode{
						value: 9,
						right: &TreeNode{
							value: 13,
						},
					},
					right: &TreeNode{
						value: 11,
					},
				},
			}},
			want: []int{8, 10, 11, 13},
		},
		{
			name: "calculate right view of the tree 2",
			args: args{tree: &TreeNode{
				value: 8,
				left: &TreeNode{
					value: 6,
					left: &TreeNode{
						value: 5,
					},
					right: &TreeNode{
						value: 7,
					},
				},
				right: &TreeNode{
					value: 10,
				},
			}},
			want: []int{8, 10, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightViewOfTreeNode(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightViewOfTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
