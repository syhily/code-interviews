package chapter3

import "testing"

func Test_countAllPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find all palindrome",
			args: args{s: "abcba"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAllPalindrome(tt.args.s); got != tt.want {
				t.Errorf("countAllPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
