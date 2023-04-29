package chapter07

import "testing"

func Test_recentAverage_ping(t *testing.T) {
	r := newRecentAverage(3000)
	tests := []struct {
		args int
		want int
	}{
		{
			args: 1,
			want: 1,
		},
		{
			args: 10,
			want: 2,
		},
		{
			args: 3001,
			want: 3,
		},
		{
			args: 3002,
			want: 3,
		},
	}
	t.Run("Ping with a fixed range of 3000", func(t *testing.T) {
		for _, tt := range tests {
			if got := r.ping(tt.args); got != tt.want {
				t.Errorf("ping() = %v, want %v", got, tt.want)
			}
		}
	})
}
