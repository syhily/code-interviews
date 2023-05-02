package chapter04

import (
	"reflect"
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_reverseListNode(t *testing.T) {
	type args struct {
		l *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "Test for nil list node",
			args: args{l: nil},
			want: nil,
		},
		{
			name: "Test for single list node",
			args: args{l: common.NewListNode(1)},
			want: common.NewListNode(1),
		},
		{
			name: "Test for 1, 2, 3, 4, 5, 6, 7",
			args: args{l: common.NewListNode(1, 2, 3, 4, 5, 6, 7)},
			want: common.NewListNode(7, 6, 5, 4, 3, 2, 1),
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
