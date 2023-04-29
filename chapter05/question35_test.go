package chapter05

import "testing"

func Test_minimalTimeDelta(t *testing.T) {
	type args struct {
		t []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "same time should return 0 as the minimal delta",
			args: args{t: []string{"12:12", "12:12", "21:32"}},
			want: 0,
		},
		{
			name: "Consider the 00:00 as the next day",
			args: args{t: []string{"23:50", "23:59", "00:00"}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimalTimeDelta(tt.args.t...); got != tt.want {
				t.Errorf("minimalTimeDelta() = %v, want %v", got, tt.want)
			}
		})
	}
}
