package chapter08

import (
	"reflect"
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_findTheNextNode(t *testing.T) {
	type args struct {
		node   *common.TreeNode[int]
		number int
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode[int]
	}{
		{
			name: "find the next node from the binary search tree",
			args: args{
				node: &common.TreeNode[int]{
					Value: 8,
					Left: &common.TreeNode[int]{
						Value: 6,
						Left: &common.TreeNode[int]{
							Value: 5,
						},
						Right: &common.TreeNode[int]{
							Value: 7,
						},
					},
					Right: &common.TreeNode[int]{
						Value: 10,
						Left: &common.TreeNode[int]{
							Value: 9,
						},
						Right: &common.TreeNode[int]{
							Value: 11,
						},
					},
				},
				number: 8,
			},
			want: &common.TreeNode[int]{
				Value: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheNextNode(tt.args.node, tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTheNextNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
