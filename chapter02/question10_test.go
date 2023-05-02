package chapter02

import "testing"

func Test_sumEqualKArrays(t *testing.T) {
	type args struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find the k sum arrays",
			args: args{
				k:    2,
				nums: []int{1, 1, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumEqualKArrays(tt.args.k, tt.args.nums...); got != tt.want {
				t.Errorf("sumEqualKArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
