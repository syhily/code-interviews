package chapter04

import (
	"reflect"
	"testing"
)

func Test_findStartNodeInCircle(t *testing.T) {
	l1, s1 := NewCircleNode(4, 1, 2, 3, 4, 5, 6)
	l2, s2 := NewCircleNode(1, 1, 2, 3, 4, 5, 6)

	type args struct {
		l *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "No start node in pure list node",
			args: args{l: New(1, 3, 5, 6, 7)},
			want: nil,
		},
		{
			name: "Find start node 1, 2, 3, 4, 5, 6 -> 4",
			args: args{l: l1},
			want: s1,
		},
		{
			name: "Find start node 1, 2, 3, 4, 5, 6 -> 1",
			args: args{l: l2},
			want: s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStartNodeInCircle(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findStartNodeInCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
