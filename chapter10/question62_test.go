package chapter10

import (
	"testing"
)

func Test_newTrie(t *testing.T) {
	type want struct {
		words  []string
		prefix []string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "word fuzzy tests",
			want: want{
				words:  []string{"all", "gaul", "is", "divided", "into", "three", "parts", "one", "of", "which"},
				prefix: []string{"al", "gau", "i", "divide", "int", "thre", "part", "on", "o"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := newTrie()
			for _, word := range tt.want.words {
				if tree.search(word) {
					t.Errorf("Trie shouldn't contain word %v", word)
				}
				tree.insert(word)
				if !tree.search(word) {
					t.Errorf("Trie should contain word %v", word)
				}
			}

			for _, prefix := range tt.want.prefix {
				if !tree.startWith(prefix) {
					t.Errorf("Trie should contain prefix %v", prefix)
				}
			}
		})
	}
}
