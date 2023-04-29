package chapter07

import (
	"reflect"
	"testing"
)

func Test_maxValuesForTreeNode(t *testing.T) {
	type args struct {
		tree *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "calculate the max values for each layers of the tree",
			args: args{tree: &TreeNode{
				value: 3,
				left: &TreeNode{
					value: 4,
					left: &TreeNode{
						value: 5,
					},
					right: &TreeNode{
						value: 1,
					},
				},
				right: &TreeNode{
					value: 2,
					right: &TreeNode{
						value: 9,
					},
				},
			}},
			want: []int{3, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValuesForTreeNode(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxValuesForTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
