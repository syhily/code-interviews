package chapter12

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "quick sort on an array",
			args: args{numbers: []int{4, 1, 5, 6, 2, 7, 8, 3}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.numbers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTheKth(t *testing.T) {
	type args struct {
		k       int
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find the kth number",
			args: args{
				k:       3,
				numbers: []int{3, 1, 2, 4, 5, 5, 6},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheKth(tt.args.k, tt.args.numbers...); got != tt.want {
				t.Errorf("findTheKth() = %v, want %v", got, tt.want)
			}
		})
	}
}
