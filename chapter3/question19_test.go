package chapter3

import "testing"

func Test_validateIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is not palindrome",
			args: args{s: "abcdefg"},
			want: false,
		},
		{
			name: "is approaching palindrome",
			args: args{s: "abcdefdcba"},
			want: true,
		},
		{
			name: "is approaching palindrome",
			args: args{s: "abca"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateIsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validateIsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
