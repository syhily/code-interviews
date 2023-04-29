package chapter07

import "testing"

func Test_leftValueFromLowestLayer(t *testing.T) {
	type args struct {
		tree *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "calculate the button left values from the tree",
			args: args{tree: &TreeNode{
				value: 8,
				left: &TreeNode{
					value: 6,
				},
				right: &TreeNode{
					value: 10,
					left: &TreeNode{
						value: 9,
						right: &TreeNode{
							value: 13,
						},
					},
					right: &TreeNode{
						value: 11,
					},
				},
			}},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftValueFromLowestLayer(tt.args.tree); got != tt.want {
				t.Errorf("leftValueFromLowestLayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
