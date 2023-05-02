package chapter02

import "testing"

func Test_multipleNumberBelowK(t *testing.T) {
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
			name: "Multiple numbers",
			args: args{
				target: 100,
				nums:   []int{10, 5, 2, 6},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multipleNumberBelowK(tt.args.target, tt.args.nums...); got != tt.want {
				t.Errorf("multipleNumberBelowK() = %v, want %v", got, tt.want)
			}
		})
	}
}
