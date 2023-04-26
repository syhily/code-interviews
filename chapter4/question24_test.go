package chapter4

import (
	"reflect"
	"testing"
)

func Test_reverseListNode(t *testing.T) {
	type args struct {
		l *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test for nil list node",
			args: args{l: nil},
			want: nil,
		},
		{
			name: "Test for single list node",
			args: args{l: New(1)},
			want: New(1),
		},
		{
			name: "Test for 1, 2, 3, 4, 5, 6, 7",
			args: args{l: New(1, 2, 3, 4, 5, 6, 7)},
			want: New(7, 6, 5, 4, 3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListNode(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
