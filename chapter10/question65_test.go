package chapter10

import "testing"

func Test_shortestEncodingLength(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "encode the array: time, me, bell",
			args: args{words: []string{"time", "me", "bell"}},
			want: 10,
		},
		{
			name: "encode the array: at, bat, cat",
			args: args{words: []string{"at", "bat", "cat"}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestEncodingLength(tt.args.words...); got != tt.want {
				t.Errorf("shortestEncodingLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
