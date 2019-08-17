package p003x

import (
    "reflect"
    "testing"
)

func Test_nextPermutation(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            args: args{[]int{1, 2, 3}},
            want: []int{1, 3, 2},
        },
        {
            args: args{[]int{1, 3, 2}},
            want: []int{2, 1, 3},
        },
        {
            args: args{[]int{3, 2, 1}},
            want: []int{1, 2, 3},
        },
        {
            args: args{[]int{1, 1, 5}},
            want: []int{1, 5, 1},
        },
        {
            args: args{[]int{1, 2, 4, 5, 3}},
            want: []int{1, 2, 5, 3, 4},
        },
        {
            args: args{[]int{1, 2, 7, 4, 3, 6, 5}},
            want: []int{1, 2, 7, 4, 5, 3, 6},
        },
        {
            args: args{[]int{1, 2, 7, 4, 6, 5, 3}},
            want: []int{1, 2, 7, 5, 3, 4, 6},
        },
        {
            args: args{[]int{}},
            want: []int{},
        },
        {
            args: args{[]int{1}},
            want: []int{1},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            nextPermutation(tt.args.nums)
            if !reflect.DeepEqual(tt.args.nums, tt.want) {
                t.Errorf("nextPermutation() = %v, want %v", tt.args.nums, tt.want)
            }
        })
    }
}
