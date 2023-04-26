package chapter4

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		node1 *ListNode
		node2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
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
				node2: New(1, 2, 3, 4, 5),
			},
			want: New(1, 2, 3, 4, 5),
		},
		{
			name: "Add two list with the same length",
			args: args{
				node1: New(1, 2, 3, 4),
				node2: New(2, 3, 4, 8),
			},
			want: New(3, 5, 8, 2),
		},
		{
			name: "Add two list with different length",
			args: args{
				node1: New(1, 4, 6, 8),
				node2: New(4, 6, 8, 9, 0, 1),
			},
			want: New(4, 7, 0, 3, 6, 9),
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
