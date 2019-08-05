package p001x

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{[]int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			args: args{[]int{-4, -1, -1, 0, 1, 2}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			args: args{[]int{2, 1, 0, -1, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			args: args{[]int{}},
			want: [][]int{},
		},
		{
			args: args{[]int{0, 0}},
			want: [][]int{},
		},
		{
			args: args{[]int{-4, -1, -4, 0, 2, -2, -4, -3, 2, -3, 2, 3, 3, -4}},
			want: [][]int{{-4, 2, 2}, {-3, 0, 3}, {-2, -1, 3}, {-2, 0, 2}},
		},
		{
			args: args{[]int{1, 2, -2, -1}},
			want: [][]int{},
		},
		{
			args: args{[]int{0,1,1}},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
