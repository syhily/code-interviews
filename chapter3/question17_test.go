package chapter3

import "testing"

func Test_findTheShortestSubstring(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string will return empty result",
			args: args{
				s: "abcdef",
				t: "",
			},
			want: "",
		},
		{
			name: "Find shortest substring",
			args: args{
				s: "ADDBANCAD",
				t: "ABC",
			},
			want: "BANC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheShortestSubstring(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheShortestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
