package chapter08

import (
	"reflect"
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_flattenTreeNode(t *testing.T) {
	type args struct {
		node *common.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode[int]
	}{
		{
			name: "flatten binary search tree",
			args: args{&common.TreeNode[int]{
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
					Value: 5,
					Right: &common.TreeNode[int]{
						Value: 6,
					},
				},
			}},
			want: &common.TreeNode[int]{
				Value: 1,
				Right: &common.TreeNode[int]{
					Value: 2,
					Right: &common.TreeNode[int]{
						Value: 3,
						Right: &common.TreeNode[int]{
							Value: 4,
							Right: &common.TreeNode[int]{
								Value: 5,
								Right: &common.TreeNode[int]{
									Value: 6,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flattenTreeNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
