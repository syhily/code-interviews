package chapter07

import (
	"reflect"
	"testing"
)

func Test_cbtInserter_insert(t *testing.T) {
	t1 := &TreeNode{
		value: 1,
		left: &TreeNode{
			value: 2,
			left: &TreeNode{
				value: 4,
			},
			right: &TreeNode{
				value: 5,
			},
		},
		right: &TreeNode{
			value: 3,
			left: &TreeNode{
				value: 6,
			},
		},
	}
	r1 := &TreeNode{
		value: 3,
		left: &TreeNode{
			value: 6,
		},
		right: &TreeNode{
			value: 7,
		},
	}
	r2 := &TreeNode{
		value: 4,
		left: &TreeNode{
			value: 8,
		},
	}
	r3 := &TreeNode{
		value: 4,
		left: &TreeNode{
			value: 8,
		},
		right: &TreeNode{
			value: 9,
		},
	}

	c := newCBTInserter(t1)
	tests := []struct {
		args int
		want *TreeNode
	}{
		{
			args: 7,
			want: r1,
		},
		{
			args: 8,
			want: r2,
		},
		{
			args: 9,
			want: r3,
		},
	}
	t.Run("Insert CBT Tree", func(t *testing.T) {
		for _, tt := range tests {
			if got := c.insert(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		}
	})
}
