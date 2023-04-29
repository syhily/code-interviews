package chapter04

import (
	"reflect"
	"testing"
)

func Test_findIntersectionNode(t *testing.T) {
	s := New(1, 3, 5, 7, 9)
	l1, l2, l3 := NewJoinNode([]int{1, 2, 3, 4}, []int{5, 6, 7, 8, 9, 10}, []int{11, 12, 13, 14, 15, 16})

	type args struct {
		f *ListNode
		s *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Two different list node",
			args: args{f: New(1, 2, 3, 4, 5, 6), s: New(1, 2, 3, 4, 5, 6, 7)},
			want: nil,
		},
		{
			name: "Two same list node",
			args: args{f: s, s: s},
			want: s,
		},
		{
			name: "Two joined list node",
			args: args{f: l1, s: l2},
			want: l3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIntersectionNode(tt.args.f, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
