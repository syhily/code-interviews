package chapter5

import (
	"reflect"
	"testing"
)

func Test_randomSet_insert(t *testing.T) {
	type fields struct {
		values   []int
		position map[int]int
	}
	type args struct {
		value int
	}
	type want struct {
		res bool
		r   *randomSet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "Insert value into an empty set",
			fields: fields{
				values:   []int{},
				position: map[int]int{},
			},
			args: args{value: 1},
			want: want{
				res: true,
				r: &randomSet{
					values:   []int{1},
					position: map[int]int{1: 0},
				},
			},
		},
		{
			name: "Insert a same value into a set",
			fields: fields{
				values:   []int{1},
				position: map[int]int{1: 0},
			},
			args: args{value: 1},
			want: want{
				res: false,
				r: &randomSet{
					values:   []int{1},
					position: map[int]int{1: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &randomSet{
				values:   tt.fields.values,
				position: tt.fields.position,
			}
			if got := r.insert(tt.args.value); got != tt.want.res && reflect.DeepEqual(r, tt.want.r) {
				t.Errorf("insert() = %v, want %v, randomSet = %v, want %v", got, tt.want, r, tt.want.r)
			}
		})
	}
}

func Test_randomSet_remove(t *testing.T) {
	type fields struct {
		values   []int
		position map[int]int
	}
	type args struct {
		value int
	}
	type want struct {
		res bool
		r   *randomSet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "Remove a value from an empty set",
			fields: fields{
				values:   []int{},
				position: map[int]int{},
			},
			args: args{value: 1},
			want: want{
				res: false,
				r: &randomSet{
					values:   []int{},
					position: map[int]int{},
				},
			},
		},
		{
			name: "Remove an existing value from an set",
			fields: fields{
				values:   []int{1, 2, 3, 4},
				position: map[int]int{1: 0, 2: 1, 3: 2, 4: 3},
			},
			args: args{value: 2},
			want: want{
				res: true,
				r: &randomSet{
					values:   []int{1, 4, 3},
					position: map[int]int{1: 0, 4: 1, 3: 2},
				},
			},
		},
		{
			name: "Remove a non-existed value from an set",
			fields: fields{
				values:   []int{1, 2, 3, 4},
				position: map[int]int{1: 0, 2: 1, 3: 2, 4: 3},
			},
			args: args{value: 5},
			want: want{
				res: false,
				r: &randomSet{
					values:   []int{1, 2, 3, 4},
					position: map[int]int{1: 0, 2: 1, 3: 2, 4: 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &randomSet{
				values:   tt.fields.values,
				position: tt.fields.position,
			}
			if got := r.remove(tt.args.value); got != tt.want.res && reflect.DeepEqual(r, tt.want.r) {
				t.Errorf("remove() = %v, want %v, randomSet = %v, want %v", got, tt.want, r, tt.want.r)
			}
		})
	}
}
