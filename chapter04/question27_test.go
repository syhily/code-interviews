package chapter04

import "testing"

func Test_isPalindromeListNode(t *testing.T) {
	type args struct {
		l *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil list assertion",
			args: args{l: nil},
			want: true,
		},
		{
			name: "Single list node assertion",
			args: args{l: New(1)},
			want: true,
		},
		{
			name: "Two list node assertion: 1, 2",
			args: args{l: New(1, 2)},
			want: false,
		},
		{
			name: "Two list node assertion: 1, 1",
			args: args{l: New(1, 1)},
			want: true,
		},
		{
			name: "Even list node assertion: 1, 1, 2, 3",
			args: args{l: New(1, 1, 2, 3)},
			want: false,
		},
		{
			name: "Even list node assertion: 1, 2, 2, 1",
			args: args{l: New(1, 2, 2, 1)},
			want: true,
		},
		{
			name: "Even list node assertion: 1, 2, 3, 3, 2, 1",
			args: args{l: New(1, 2, 3, 3, 2, 1)},
			want: true,
		},
		{
			name: "Even list node assertion: 1, 2, 3, 3, 3, 1",
			args: args{l: New(1, 2, 3, 3, 3, 1)},
			want: false,
		},
		{
			name: "Odd list node assertion: 1, 1, 4, 2, 3",
			args: args{l: New(1, 1, 4, 2, 3)},
			want: false,
		},
		{
			name: "Odd list node assertion: 1, 2, 4, 2, 1",
			args: args{l: New(1, 2, 4, 2, 1)},
			want: true,
		},
		{
			name: "Odd list node assertion: 1, 2, 3, 5, 3, 2, 1",
			args: args{l: New(1, 2, 3, 5, 3, 2, 1)},
			want: true,
		},
		{
			name: "Odd list node assertion: 1, 2, 3, 5, 3, 3, 1",
			args: args{l: New(1, 2, 3, 5, 3, 3, 1)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeListNode(tt.args.l); got != tt.want {
				t.Errorf("isPalindromeListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
