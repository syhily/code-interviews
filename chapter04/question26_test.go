package chapter04

import (
	"reflect"
	"testing"
)

func Test_reorderListNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Reorder nil list node",
			args: args{node: nil},
			want: nil,
		},
		{
			name: "Reorder single list node",
			args: args{node: New(1)},
			want: New(1),
		},
		{
			name: "Reorder odd list node",
			args: args{node: New(1, 3, 5, 7, 9, 11, 13)},
			want: New(1, 13, 3, 11, 5, 9, 7),
		},
		{
			name: "Reorder even list node",
			args: args{node: New(2, 4, 6, 8, 10, 12)},
			want: New(2, 12, 4, 10, 6, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderListNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
