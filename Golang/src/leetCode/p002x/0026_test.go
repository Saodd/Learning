package p002x

import (
    "reflect"
    "testing"
)

func Test_removeDuplicates(t *testing.T) {
    type args struct {
        nums []int
    }
    type wants struct {
        nums []int
        num  int
    }
    tests := []struct {
        name string
        args args
        want wants
    }{
        {
            args: args{[]int{1, 1, 2}},
            want: wants{[]int{1, 2}, 2},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := removeDuplicates(tt.args.nums);
            if !reflect.DeepEqual(tt.args.nums[:got], tt.want.nums[:got]) || got != tt.want.num {
                t.Errorf("removeDuplicates(%v) = %v, want %v, %v", tt.args.nums, got, tt.want.nums, tt.want.num)
            }
        })
    }
}
