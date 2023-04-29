package chapter03

import "testing"

func Test_findInclusionWord(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s1's length is greater than s2's length",
			args: args{s1: "abcdefg", s2: "abcdef"},
			want: false,
		},
		{
			name: "s1 isn't included in s2",
			args: args{s1: "abcdefg", s2: "hijklmnop"},
			want: false,
		},
		{
			name: "s1 is included in s2",
			args: args{s1: "abcdefg", s2: "hijgfedcbacd"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInclusionWord(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("findInclusionWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
