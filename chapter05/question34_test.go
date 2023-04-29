package chapter05

import "testing"

func Test_isSortedWords(t *testing.T) {
	type args struct {
		alphabet string
		words    []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same word prefix with sorted end",
			args: args{
				alphabet: "abcdefghijklmnopqrstuvwxyz",
				words:    []string{"a", "ab", "abc", "abcd"},
			},
			want: true,
		},
		{
			name: "same word prefix with unsorted end",
			args: args{
				alphabet: "abcdefghijklmnopqrstuvwxyz",
				words:    []string{"a", "abcd", "abc", "abc"},
			},
			want: false,
		},
		{
			name: "unsorted words",
			args: args{
				alphabet: "abcdefghijklmnopqrstuvwxyz",
				words:    []string{"db", "abcd", "abc", "abc"},
			},
			want: false,
		},
		{
			name: "sorted words",
			args: args{
				alphabet: "zyxwvutsrqponmlkjihgfedcba",
				words:    []string{"offer", "is", "coming"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSortedWords(tt.args.alphabet, tt.args.words...); got != tt.want {
				t.Errorf("isSortedWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
