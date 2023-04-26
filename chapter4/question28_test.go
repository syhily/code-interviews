package chapter4

import (
	"reflect"
	"testing"
)

func Test_flattenNode(t *testing.T) {
	node1 := createNode(
		CreateNode{value: 1, next: 2},
		CreateNode{value: 2, next: 3, child: 5, prev: 1},
		CreateNode{value: 3, next: 4, prev: 2},
		CreateNode{value: 4, prev: 3},
		CreateNode{value: 5, next: 6},
		CreateNode{value: 6, next: 7, child: 8, prev: 5},
		CreateNode{value: 7, prev: 6},
		CreateNode{value: 8, next: 9},
		CreateNode{value: 9, prev: 8},
	)
	node2 := createNode(
		CreateNode{value: 1, next: 2},
		CreateNode{value: 2, next: 5, prev: 1},
		CreateNode{value: 5, next: 6, prev: 2},
		CreateNode{value: 6, next: 8, prev: 5},
		CreateNode{value: 8, next: 9, prev: 6},
		CreateNode{value: 9, next: 7, prev: 8},
		CreateNode{value: 7, next: 3, prev: 9},
		CreateNode{value: 3, next: 4, prev: 7},
		CreateNode{value: 4, prev: 3},
	)

	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Flatten nodes",
			args: args{node: node1},
			want: node2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flattenNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
