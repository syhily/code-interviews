package chapter04

import (
	"reflect"
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_removeNthFromTheEnd(t *testing.T) {
	type args struct {
		l *common.ListNode
		n int
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "Delete first node",
			args: args{
				l: common.NewListNode(1, 2, 3, 4, 5, 6),
				n: 6,
			},
			want: common.NewListNode(2, 3, 4, 5, 6),
		},
		{
			name: "Delete last node",
			args: args{
				l: common.NewListNode(1, 2, 3, 4, 5, 6),
				n: 1,
			},
			want: common.NewListNode(1, 2, 3, 4, 5),
		},
		{
			name: "Delete nth node",
			args: args{
				l: common.NewListNode(1, 2, 3, 4, 5, 6),
				n: 3,
			},
			want: common.NewListNode(1, 2, 3, 5, 6),
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
