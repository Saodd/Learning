package p004x

import (
    "reflect"
    "testing"
)

func Test_permuteUnique(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want [][]int
    }{
        {
            name: "示例",
            args: args{[]int{1, 1, 2}},
            want: [][]int{
                {1, 1, 2}, {1, 2, 1}, {2, 1, 1},
            },
        },
        {
            name: "空数组",
            args: args{[]int{}},
            want: [][]int{
            },
        },
        {
            name: "大数重复",
            args: args{[]int{1, 2, 2}},
            want: [][]int{
                {1, 2, 2}, {2, 1, 2}, {2, 2, 1},
            },
        },
        {
            name: "不重复",
            args: args{[]int{1, 2, 3}},
            want: [][]int{
                {1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2},
            },
        },
        {
            name: "多个重复值",
            args: args{[]int{0,1,0,0,9}},
            want: [][]int{
                {0,0,0,1,9},{0,0,0,9,1},{0,0,1,0,9},{0,0,1,9,0},{0,0,9,0,1},{0,0,9,1,0},
                {0,1,0,0,9},{0,1,0,9,0},{0,1,9,0,0},{0,9,0,0,1},{0,9,0,1,0},{0,9,1,0,0},
                {1,0,0,0,9},{1,0,0,9,0},{1,0,9,0,0},{1,9,0,0,0},{9,0,0,0,1},{9,0,0,1,0},
                {9,0,1,0,0},{9,1,0,0,0},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
            }
        })
    }
}
