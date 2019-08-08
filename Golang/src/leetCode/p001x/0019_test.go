package p001x

import (
	"fmt"
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

func TestSliceToChain(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{[]int{}},
			want: nil,
		},
		{
			args: args{[]int{5, 6, 7}},
			want: &ListNode{5, &ListNode{6, &ListNode{7, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToChain(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToChain() = %v, want %v", got, tt.want)
			}
		})
		//    head := SliceToChain([]int{5,6,7})
		//    for head!=nil{
		//        fmt.Println(head.Val)
		//        head = head.Next
		//    }
	}
}

func TestChainToSlice(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{&ListNode{5, &ListNode{6, &ListNode{7, nil}}}},
			want: []int{5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChainToSlice(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChainToSlice() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}
