package chapter06

import "testing"

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "calculate: 2 1 3 * +",
			args: args{tokens: []string{"2", "1", "3", "*", "+"}},
			want: 5,
		},
		{
			name: "calculate: 2 1 3 + * 4 / 2 -",
			args: args{tokens: []string{"2", "1", "3", "+", "*", "4", "/", "2", "-"}},
			want: 0,
		},
		{
			name: "calculate single token: 4",
			args: args{tokens: []string{"4"}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens...); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
