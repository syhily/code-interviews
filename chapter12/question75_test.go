package chapter12

import (
	"reflect"
	"testing"
)

func Test_sortByGivenIndex(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort the arr1 by the arr2",
			args: args{
				arr1: []int{2, 3, 3, 7, 3, 9, 2, 1, 7, 2},
				arr2: []int{3, 2, 1},
			},
			want: []int{3, 3, 3, 2, 2, 2, 1, 7, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByGivenIndex(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByGivenIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
