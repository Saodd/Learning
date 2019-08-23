package p004x

import (
    "reflect"
    "testing"
)

func Test_combinationSum2(t *testing.T) {
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
            name: "示例",
            args: args{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
            want: [][]int{
                {1, 2, 5}, {1, 7}, {1, 1, 6}, {2, 6},
            },
        },
        {
            name: "示例",
            args: args{[]int{2, 5, 2, 1, 2}, 5},
            want: [][]int{
                {1, 2, 2}, {5},
            },
        },
        {
            name: "无解",
            args: args{[]int{2, 5, 2, 1, 2}, 0},
            want: [][]int{
            },
        },
        {
            name: "多个重复值",
            args: args{[]int{1, 1, 1, 1, 2, 2, 3, 3}, 5},
            want: [][]int{
                {1, 2, 2}, {1, 1, 3}, {1, 1, 1, 2}, {2, 3},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
            }
        })
    }
}
