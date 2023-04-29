package chapter04

import (
	"reflect"
	"testing"
)

func Test_removeNthFromTheEnd(t *testing.T) {
	type args struct {
		l *ListNode
		n int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Delete first node",
			args: args{
				l: New(1, 2, 3, 4, 5, 6),
				n: 6,
			},
			want: New(2, 3, 4, 5, 6),
		},
		{
			name: "Delete last node",
			args: args{
				l: New(1, 2, 3, 4, 5, 6),
				n: 1,
			},
			want: New(1, 2, 3, 4, 5),
		},
		{
			name: "Delete nth node",
			args: args{
				l: New(1, 2, 3, 4, 5, 6),
				n: 3,
			},
			want: New(1, 2, 3, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromTheEnd(tt.args.l, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromTheEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
