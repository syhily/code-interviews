package chapter08

import (
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_twoSum(t *testing.T) {
	tree := &common.TreeNode[int]{
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
	}

	type args struct {
		node *common.TreeNode[int]
		num  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "find two sum",
			args: args{
				node: tree,
				num:  12,
			},
			want: true,
		},
		{
			name: "couldn't find two sum",
			args: args{
				node: tree,
				num:  22,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.node, tt.args.num); got != tt.want {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
