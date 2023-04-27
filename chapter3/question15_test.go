package chapter3

import (
	"reflect"
	"testing"
)

func Test_findAllStartIndexOfInclusionWords(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Find the inclusion words' start indexes",
			args: args{
				s1: "cbadabacg",
				s2: "abc",
			},
			want: []int{0, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllStartIndexOfInclusionWords(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllStartIndexOfInclusionWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
