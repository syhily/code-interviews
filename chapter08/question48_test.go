package chapter08

import (
	"reflect"
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_serializeAndDeserializeTreeNode(t *testing.T) {
	type args struct {
		node *common.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "serialize tree node 1",
			args: args{&common.TreeNode[int]{
				Value: 6,
				Left: &common.TreeNode[int]{
					Value: 6,
					Left: &common.TreeNode[int]{
						Value: 6,
					},
					Right: &common.TreeNode[int]{
						Value: 6,
					},
				},
				Right: &common.TreeNode[int]{
					Value: 6,
				},
			}},
			want: "6,6,6,#,#,6,#,#,6,#,#",
		},
		{
			name: "serialize tree node 1",
			args: args{&common.TreeNode[int]{
				Value: 6,
				Left: &common.TreeNode[int]{
					Value: 6,
				},
				Right: &common.TreeNode[int]{
					Value: 6,
					Left: &common.TreeNode[int]{
						Value: 6,
					},
					Right: &common.TreeNode[int]{
						Value: 6,
					},
				},
			}},
			want: "6,6,#,#,6,6,#,#,6,#,#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serializeTreeNode(tt.args.node); got != tt.want {
				t.Errorf("serializeTreeNode() = %v, want %v", got, tt.want)
			}
			if got := deserializeTreeNode(tt.want); !reflect.DeepEqual(got, tt.args.node) {
				t.Errorf("deserializeTreeNode() = %v, want %v", got, tt.args.node)
			}
		})
	}
}
