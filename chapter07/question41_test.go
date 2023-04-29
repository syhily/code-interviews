package chapter07

import (
	"testing"
)

func Test_movingAverage_next(t *testing.T) {
	moves := []MovingAverage{
		newMovingAverageSlice(5),
		newMovingAverageQueue(5),
	}
	tests := []struct {
		args int
		want float64
	}{
		{
			args: 10,
			want: 10.0,
		},
		{
			args: 8,
			want: 9.0,
		},
		{
			args: 6,
			want: 8.0,
		},
		{
			args: 4,
			want: 7.0,
		},
		{
			args: 2,
			want: 6.0,
		},
		{
			args: 10,
			want: 6.0,
		},
		{
			args: 8,
			want: 6.0,
		},
		{
			args: 6,
			want: 6.0,
		},
		{
			args: 4,
			want: 6.0,
		},
		{
			args: 2,
			want: 6.0,
		},
		{
			args: 10,
			want: 6.0,
		},
		{
			args: 8,
			want: 6.0,
		},
		{
			args: 6,
			want: 6.0,
		},
		{
			args: 4,
			want: 6.0,
		},
		{
			args: 2,
			want: 6.0,
		},
		{
			args: 10,
			want: 6.0,
		},
		{
			args: 8,
			want: 6.0,
		},
		{
			args: 6,
			want: 6.0,
		},
		{
			args: 4,
			want: 6.0,
		},
		{
			args: 2,
			want: 6.0,
		},
	}
	t.Run("execute the next calculation", func(t *testing.T) {
		for _, move := range moves {
			for _, tt := range tests {
				if got := move.next(tt.args); got != tt.want {
					t.Errorf("next() = %v, want %v", got, tt.want)
				}
			}
		}
	})
}
