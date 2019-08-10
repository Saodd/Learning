package p001x

import (
    "fmt"
    . "leetCode/utils"
    "reflect"
    "testing"
)

func Test_removeNthFromEnd(t *testing.T) {
    {
        arg := SliceToChain([]int{1, 2, 3, 4, 5})
        want := []int{1, 2, 3, 5}
        got := ChainToSlice(removeNthFromEnd(arg, 2))
        fmt.Println(reflect.DeepEqual(want, got), want, got)
    }
    {
        arg := SliceToChain([]int{1})
        want := []int{}
        got := ChainToSlice(removeNthFromEnd(arg, 1))
        fmt.Println(reflect.DeepEqual(want, got), want, got)
    }
    {
        arg := SliceToChain([]int{1, 2})
        want := []int{1}
        got := ChainToSlice(removeNthFromEnd(arg, 1))
        fmt.Println(reflect.DeepEqual(want, got), want, got)
    }
}

