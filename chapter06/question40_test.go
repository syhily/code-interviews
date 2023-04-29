package chapter06

import "testing"

func Test_maximumRectangle(t *testing.T) {
	type args struct {
		matrix [][]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "calculate the maximum matrix",
			args: args{matrix: [][]bool{
				{true, false, true, false, false},
				{false, false, true, true, true},
				{true, true, true, true, true},
				{true, false, false, true, false},
			}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximumRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
