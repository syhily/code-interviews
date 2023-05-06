package chapter08

import (
	"reflect"
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_removeZeroNode(t *testing.T) {
	type args struct {
		node *common.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode[int]
	}{
		{
			name: "nil tree node will return ni",
			args: args{node: nil},
			want: nil,
		},
		{
			name: "remove all zero node",
			args: args{node: &common.TreeNode[int]{
				Value: 1,
				Left: &common.TreeNode[int]{
					Value: 0,
					Left: &common.TreeNode[int]{
						Value: 0,
					},
					Right: &common.TreeNode[int]{
						Value: 0,
					},
				},
				Right: &common.TreeNode[int]{
					Value: 0,
					Left: &common.TreeNode[int]{
						Value: 0,
					},
					Right: &common.TreeNode[int]{
						Value: 1,
					},
				},
			}},
			want: &common.TreeNode[int]{
				Value: 1,
				Right: &common.TreeNode[int]{
					Value: 0,
					Right: &common.TreeNode[int]{
						Value: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeZeroNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeZeroNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
