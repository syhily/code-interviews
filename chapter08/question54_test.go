package chapter08

import (
	"reflect"
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_summarizeTheTreeNode(t *testing.T) {
	type args struct {
		node *common.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode[int]
	}{
		{
			name: "summarize the tree node",
			args: args{node: &common.TreeNode[int]{
				Value: 4,
				Left: &common.TreeNode[int]{
					Value: 2,
					Left: &common.TreeNode[int]{
						Value: 1,
					},
					Right: &common.TreeNode[int]{
						Value: 3,
					},
				},
				Right: &common.TreeNode[int]{
					Value: 6,
					Left: &common.TreeNode[int]{
						Value: 5,
					},
					Right: &common.TreeNode[int]{
						Value: 7,
					},
				},
			}},
			want: &common.TreeNode[int]{
				Value: 22,
				Left: &common.TreeNode[int]{
					Value: 27,
					Left: &common.TreeNode[int]{
						Value: 28,
					},
					Right: &common.TreeNode[int]{
						Value: 25,
					},
				},
				Right: &common.TreeNode[int]{
					Value: 13,
					Left: &common.TreeNode[int]{
						Value: 18,
					},
					Right: &common.TreeNode[int]{
						Value: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summarizeTheTreeNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summarizeTheTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
