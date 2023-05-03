package chapter02

import "testing"

func TestMatrix_calculate(t *testing.T) {
	type fields struct {
		matrix [][]int
	}
	type args struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "calculate the matrix",
			fields: fields{matrix: [][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			}},
			args: args{
				x1: 2,
				y1: 1,
				x2: 4,
				y2: 3,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMatrix(tt.fields.matrix)
			if got := m.calculate(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
