package chapter12

import (
	"reflect"
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_mergeKListNodes(t *testing.T) {
	type args struct {
		nodes []*common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "merge k sort lists",
			args: args{[]*common.ListNode{
				common.New(1, 7, 13, 19),
				common.New(2, 8, 14, 20),
				common.New(3, 9, 15, 21),
				common.New(4, 10, 16, 22),
				common.New(5, 11, 17, 23),
				common.New(6, 12, 18, 24),
			}},
			want: common.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKListNodes(tt.args.nodes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKListNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
