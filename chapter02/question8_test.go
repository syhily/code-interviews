package chapter02

import "testing"

func Test_shortestArray(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "shortest sub array 1",
			args: args{
				target: 7,
				nums:   []int{5, 1, 4, 3},
			},
			want: 2,
		},
		{
			name: "shortest sub array 2",
			args: args{
				target: 7,
				nums:   []int{5, 1, 4, 7},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestArray(tt.args.target, tt.args.nums...); got != tt.want {
				t.Errorf("shortestArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
