package chapter05

import "testing"

func Test_isAnagram(t *testing.T) {
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
			name: "same word should return false",
			args: args{
				s1: "same",
				s2: "same",
			},
			want: false,
		},
		{
			name: "two words with different length should return false",
			args: args{
				s1: "abcd",
				s2: "abcde",
			},
			want: false,
		},
		{
			name: "two words with same length by isn't anagram should return false",
			args: args{
				s1: "abcd",
				s2: "ecda",
			},
			want: false,
		},
		{
			name: "two anagram words should return true",
			args: args{
				s1: "abcd",
				s2: "dcba",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
