package p002x

import (
    . "leetCode/utils"
    "reflect"
    "testing"
)

func Test_mergeTwoLists(t *testing.T) {
    type args struct {
        l1 *ListNode
        l2 *ListNode
    }
    tests := []struct {
        name string
        args args
        want *ListNode
    }{
        {
            args: args{SliceToChain([]int{1, 2, 3}), SliceToChain([]int{4, 5, 6})},
            want: SliceToChain([]int{1, 2, 3, 4, 5, 6}),
        },
        {
            args: args{SliceToChain([]int{1, 5, 6}), SliceToChain([]int{2, 3, 4})},
            want: SliceToChain([]int{1, 2, 3, 4, 5, 6}),
        },
        {
            args: args{SliceToChain([]int{1, 2, 3}), SliceToChain([]int{})},
            want: SliceToChain([]int{1, 2, 3}),
        },
        {
            args: args{SliceToChain([]int{}), SliceToChain([]int{})},
            want: SliceToChain([]int{}),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(ChainToSlice(got), ChainToSlice(tt.want)) {
                t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
            }
        })
    }
}
