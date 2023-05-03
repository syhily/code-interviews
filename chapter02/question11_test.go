package chapter02

import "testing"

func Test_findTheMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max length array",
			args: args{nums: []int{1, 0, 1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheMaxLength(tt.args.nums...); got != tt.want {
				t.Errorf("findTheMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
