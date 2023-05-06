package chapter08

import (
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_countTheRoutine(t *testing.T) {
	type args struct {
		node    *common.TreeNode[int]
		desired int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "count the 8 from the treenode",
			args: args{
				node: &common.TreeNode[int]{
					Value: 5,
					Left: &common.TreeNode[int]{
						Value: 2,
						Left: &common.TreeNode[int]{
							Value: 1,
						},
						Right: &common.TreeNode[int]{
							Value: 6,
						},
					},
					Right: &common.TreeNode[int]{
						Value: 4,
						Left: &common.TreeNode[int]{
							Value: 3,
						},
						Right: &common.TreeNode[int]{
							Value: 7,
						},
					},
				},
				desired: 8,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTheRoutine(tt.args.node, tt.args.desired); got != tt.want {
				t.Errorf("countTheRoutine() = %v, want %v", got, tt.want)
			}
		})
	}
}
