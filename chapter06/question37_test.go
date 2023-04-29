package chapter06

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty asteroids will return empty list",
			args: args{asteroids: []int{}},
			want: []int{},
		},
		{
			name: "single asteroid will return list will single item",
			args: args{asteroids: []int{3}},
			want: []int{3},
		},
		{
			name: "asteroids [4, 5, -6, 4, 8, -5]",
			args: args{asteroids: []int{4, 5, -6, 4, 8, -5}},
			want: []int{-6, 4, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollision(tt.args.asteroids...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
