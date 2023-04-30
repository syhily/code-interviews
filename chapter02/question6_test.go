package chapter02

import "testing"

func Test_twoSumInArray(t *testing.T) {
	type args struct {
		sum    int
		arrays []int
	}
	tests := []struct {
		name  string
		args  args
		wantL int
		wantR int
	}{
		{
			name: "Find two indexes from array",
			args: args{
				sum:    8,
				arrays: []int{1, 2, 4, 6, 10},
			},
			wantL: 1,
			wantR: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotL, gotR := twoSum(tt.args.sum, tt.args.arrays...)
			if gotL != tt.wantL {
				t.Errorf("twoSum() gotL = %v, want %v", gotL, tt.wantL)
			}
			if gotR != tt.wantR {
				t.Errorf("twoSum() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
