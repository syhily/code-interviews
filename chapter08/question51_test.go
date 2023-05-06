package chapter08

import (
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_findTheMaxRoutine(t *testing.T) {
	type args struct {
		node *common.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find the max paths",
			args: args{node: &common.TreeNode[int]{
				Value: -9,
				Left: &common.TreeNode[int]{
					Value: 4,
				},
				Right: &common.TreeNode[int]{
					Value: 20,
					Left: &common.TreeNode[int]{
						Value: 15,
					},
					Right: &common.TreeNode[int]{
						Value: 7,
						Left: &common.TreeNode[int]{
							Value: -3,
						},
					},
				},
			}},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheMaxRoutine(tt.args.node); got != tt.want {
				t.Errorf("findTheMaxRoutine() = %v, want %v", got, tt.want)
			}
		})
	}
}
