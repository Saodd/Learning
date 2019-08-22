package p003x

import (
    "reflect"
    "testing"
)

func Test_combinationSum(t *testing.T) {
    type args struct {
        candidates []int
        target     int
    }
    tests := []struct {
        name string
        args args
        want [][]int
    }{
        {
            args: args{[]int{2, 3, 6, 7}, 7},
            want: [][]int{
                {2, 2, 3}, {7},
            },
        },
        {
            args: args{[]int{2, 3, 5}, 8},
            want: [][]int{
                {2, 2, 2, 2}, {2, 3, 3}, {3, 5},
            },
        },
        {
            name: "无解",
            args: args{[]int{2, 3, 5}, 1},
            want: [][]int{
            },
        },
        {
            name: "target在中间",
            args: args{[]int{1, 4, 5}, 3},
            want: [][]int{
                {1, 1, 1},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("combinationSum() = %v, want %v", got, tt.want)
            }
        })
    }
}
