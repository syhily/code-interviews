package chapter4

import (
	"reflect"
	"testing"
)

func Test_insertNodeIntoCircle(t *testing.T) {
	n1, _ := NewCircleNode(1, 1)
	n2, _ := NewCircleNode(1, 1, 2)
	n3, _ := NewCircleNode(1, 1, 2, 3, 4, 5, 6)
	n4, _ := NewCircleNode(1, 1, 2, 3, 4, 5, 6, 7)
	n5, _ := NewCircleNode(1, 1, 2, 3, 4, 5, 8, 9)
	n6, _ := NewCircleNode(1, 1, 2, 3, 4, 5, 7, 8, 9)

	type args struct {
		node  *ListNode
		value int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Insert the value into nil list node",
			args: args{node: nil, value: 1},
			want: n1,
		},
		{
			name: "Insert the value into single list node",
			args: args{node: n1, value: 2},
			want: n2,
		},
		{
			name: "Insert the value into multiple list node at the end",
			args: args{node: n3, value: 7},
			want: n4,
		},
		{
			name: "Insert the value into multiple list node in the middle",
			args: args{node: n5, value: 7},
			want: n6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertNodeIntoCircle(tt.args.node, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertNodeIntoCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
