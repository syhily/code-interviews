package chapter10

import "testing"

func Test_magicDictionary_search(t *testing.T) {
	type args struct {
		word string
		dict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "find the magic words",
			args: args{
				word: "now",
				dict: []string{"happy", "new", "year"},
			},
			want: true,
		},
		{
			name: "find the magic words",
			args: args{
				word: "yeah",
				dict: []string{"happy", "new", "year"},
			},
			want: true,
		},
		{
			name: "couldn't find the magic words",
			args: args{
				word: "new",
				dict: []string{"happy", "new", "year"},
			},
			want: false,
		},
		{
			name: "couldn't find the magic words",
			args: args{
				word: "apple",
				dict: []string{"happy", "new", "year"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &magicDictionary{}
			m.buildDict(tt.args.dict)

			if got := m.search(tt.args.word); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
