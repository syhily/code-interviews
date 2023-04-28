package chapter3

import "testing"

func Test_lengthOfTheLongestSubString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty str will return 0",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "test string with no duplicated char",
			args: args{s: "abcdefghijk"},
			want: 11,
		},
		{
			name: "test string babcca",
			args: args{s: "babcca"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfTheLongestSubString(tt.args.s); got != tt.want {
				t.Errorf("lengthOfTheLongestSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}
