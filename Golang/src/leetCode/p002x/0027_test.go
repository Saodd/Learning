package p002x

import (
    "reflect"
    "testing"
)

func Test_removeElement(t *testing.T) {
    type args struct {
        nums []int
        val  int
    }
    type wants struct {
        nums   []int
        length int
    }
    tests := []struct {
        name string
        args args
        want wants
    }{
        {
          args: args{[]int{3, 2, 2, 3}, 3},
          want: wants{[]int{2, 2, 3, 3}, 2},
        },
        {
          args: args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
          want: wants{[]int{0, 1, 4, 0, 3, 2, 2, 2}, 5},
        },
        {
          args: args{[]int{}, 2},
          want: wants{[]int{}, 0},
        },
        {
          args: args{[]int{2}, 3},
          want: wants{[]int{2}, 1},
        },
        {
           args: args{[]int{2, 2}, 3},
           want: wants{[]int{2, 2}, 2},
        },
        {
           args: args{[]int{2}, 2},
           want: wants{[]int{2}, 0},
        },
        {
            args: args{[]int{2, 2}, 2},
            want: wants{[]int{2, 2}, 0},
        },
        {
            args: args{[]int{2, 3, 3, 3, 3, 3}, 3},
            want: wants{[]int{2, 3, 3, 3, 3, 3}, 1},
        },
        {
            args: args{[]int{3, 3, 3, 3, 3, 2}, 3},
            want: wants{[]int{2, 3, 3, 3, 3, 3}, 1},
        },
        {
            args: args{[]int{2,3}, 2},
            want: wants{[]int{3,3}, 1},
        },
    }
    for i, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := removeElement(tt.args.nums, tt.args.val);
            if !reflect.DeepEqual(tt.args.nums[:got], tt.want.nums[:got]) || got != tt.want.length {
                t.Errorf("removeDuplicates(%v, %v) = %v, want %v, %v", tests[i].args.nums,tt.args.val, got, tt.want.nums, tt.want.length)
            }
        })
    }
}
