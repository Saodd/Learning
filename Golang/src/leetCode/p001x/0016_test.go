package p001x

import (
    "fmt"
    "learnAlgo"
    "learnAlgo/tools"
    "testing"
    "time"
)

func Test_threeSumClosest(t *testing.T) {
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
            args: args{[]int{-1, 2, 1, -4}, 1},
            want: 2,
        },
        {
            args: args{[]int{0, 0, 0}, 1},
            want: 0,
        },
        {
            args: args{[]int{1, 1, -1, -1, 3}, -1},
            want: -1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
                t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_threeSumClosest_Time(t *testing.T) {
    nums := tools.Gen_ints_list(3000000)
    learnAlgo.QuickSortInt(nums)
    {
        start := time.Now()
        threeSumClosest(nums, 0)
        totalTime := time.Since(start)
        fmt.Println(totalTime.Seconds() )
    }
    {
        start := time.Now()
        threeSumClosest1(nums, 0)
        totalTime := time.Since(start)
        fmt.Println(totalTime.Seconds())
    }
}
