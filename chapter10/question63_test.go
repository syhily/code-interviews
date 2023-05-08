package chapter10

import "testing"

func Test_replaceWords(t *testing.T) {
	type args struct {
		sentence string
		dict     []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "replace words by dictionary",
			args: args{
				sentence: "the cattle was rattled by the battery",
				dict:     []string{"cat", "rat", "bat"},
			},
			want: "the cat was rat by the bat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceWords(tt.args.sentence, tt.args.dict...); got != tt.want {
				t.Errorf("replaceWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
