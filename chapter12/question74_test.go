package chapter12

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Merge the intervals",
			args: args{intervals: [][]int{
				{1, 3},
				{4, 5},
				{8, 10},
				{2, 6},
				{9, 12},
				{15, 18},
			}},
			want: [][]int{
				{1, 6},
				{8, 12},
				{15, 18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeIntervals(tt.args.intervals...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
