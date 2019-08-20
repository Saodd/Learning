package p003x

import "testing"

func Test_searchInsert(t *testing.T) {
    type args struct {
        nums   []int
        target int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "中间命中",
            args: args{[]int{1, 3, 5, 6}, 5},
            want: 2,
        },
        {
            name: "左侧命中",
            args: args{[]int{1, 3, 5, 6}, 1},
            want: 0,
        },
        {
            name: "右侧命中",
            args: args{[]int{1, 3, 5, 6}, 6},
            want: 3,
        },
        {
            name: "中间未命中",
            args: args{[]int{1, 3, 5, 6}, 2},
            want: 1,
        },
        {
            name: "右侧未命中",
            args: args{[]int{1, 3, 5, 6}, 7},
            want: 4,
        },
        {
            name: "左侧未命中",
            args: args{[]int{1, 3, 5, 6}, 0},
            want: 0,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
                t.Errorf("searchInsert() = %v, want %v", got, tt.want)
            }
        })
    }
}
