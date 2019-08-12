package p002x

import (
    . "leetCode/utils"
    "reflect"
    "testing"
)

func Test_swapPairs(t *testing.T) {
    type args struct {
        head *ListNode
    }
    tests := []struct {
        name string
        args args
        want *ListNode
    }{
        {
            args: args{SliceToChain([]int{1, 2, 3, 4})},
            want: SliceToChain([]int{2, 1, 4, 3}),
        },
        {
            args: args{SliceToChain([]int{})},
            want: SliceToChain([]int{}),
        },
        {
            args: args{SliceToChain([]int{1, 2, 3})},
            want: SliceToChain([]int{2, 1, 3}),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("swapPairs() = %v, want %v", got, tt.want)
            }
        })
    }
}
