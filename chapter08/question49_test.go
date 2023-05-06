package chapter08

import (
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_calculateNumbers(t *testing.T) {
	type args struct {
		nodes *common.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "calculate the tree nodes",
			args: args{nodes: &common.TreeNode[int]{
				Value: 3,
				Left: &common.TreeNode[int]{
					Value: 9,
					Left: &common.TreeNode[int]{
						Value: 5,
					},
					Right: &common.TreeNode[int]{
						Value: 1,
					},
				},
				Right: &common.TreeNode[int]{
					Value: 0,
					Right: &common.TreeNode[int]{
						Value: 2,
					},
				},
			}},
			want: 1088,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateNumbers(tt.args.nodes); got != tt.want {
				t.Errorf("calculateNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
