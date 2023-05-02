package chapter04

import (
	"reflect"
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		node1 *common.ListNode
		node2 *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "Add two nil list",
			args: args{
				node1: nil,
				node2: nil,
			},
			want: nil,
		},
		{
			name: "Add one nil and one list",
			args: args{
				node1: nil,
				node2: common.NewListNode(1, 2, 3, 4, 5),
			},
			want: common.NewListNode(1, 2, 3, 4, 5),
		},
		{
			name: "Add two list with the same length",
			args: args{
				node1: common.NewListNode(1, 2, 3, 4),
				node2: common.NewListNode(2, 3, 4, 8),
			},
			want: common.NewListNode(3, 5, 8, 2),
		},
		{
			name: "Add two list with different length",
			args: args{
				node1: common.NewListNode(1, 4, 6, 8),
				node2: common.NewListNode(4, 6, 8, 9, 0, 1),
			},
			want: common.NewListNode(4, 7, 0, 3, 6, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.node1, tt.args.node2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
