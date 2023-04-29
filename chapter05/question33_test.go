package chapter05

import (
	"reflect"
	"testing"
)

func Test_groupingAnagrams(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Grouping the words",
			args: args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupingAnagrams(tt.args.s...); !groupingAnagramResultIsEqual(got, tt.want) {
				t.Errorf("groupingAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func groupingAnagramResultIsEqual(got, want [][]string) bool {
d:
	for _, g := range got {
		for _, w := range want {
			if reflect.DeepEqual(g, w) {
				continue d
			}
		}

		return false
	}

	return true
}
