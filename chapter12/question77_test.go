package chapter12

import (
	"reflect"
	"testing"

	"github.com/syhily/code-interviews/common"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "merging sort on an array",
			args: args{numbers: []int{6, 5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.numbers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reclusiveMergeSortList(t *testing.T) {
	type args struct {
		node *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "merge sort the link list",
			args: args{node: common.New(6, 5, 4, 3, 2, 1)},
			want: common.New(1, 2, 3, 4, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reclusiveMergeSortList(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reclusiveMergeSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSortList(t *testing.T) {
	type args struct {
		node *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "merge sort the link list",
			args: args{node: common.New(4, 1, 5, 6, 2, 7, 8, 3)},
			want: common.New(1, 2, 3, 4, 5, 6, 7, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSortList(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
