package learnAlgo

import (
	"fmt"
	"learnAlgo/tools"
	"testing"
)

func Test_quickSortIntSelectSort(t *testing.T) {
	type args struct {
		li []int
		lo int
		hi int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{tools.Gen_ints_list(10), 0, 9},
		},
		{
			args: args{[]int{3, 2, 1}, 0, 2},
		},
		{
			args: args{[]int{1, 2, 3}, 0, 2},
		},
		{
			args: args{[]int{}, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSortIntSelectSort(tt.args.li, tt.args.lo, tt.args.hi)
			if !isSorted(tt.args.li) {
				t.Error(fmt.Sprint(tt.args.li))
			}

		})
	}
}

func TestQuickSortInt(t *testing.T) {
	type args struct {
		li []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{tools.Gen_ints_list(2000000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSortInt(tt.args.li)
			if !isSorted(tt.args.li) {
				t.Fail()
			}
		})
	}
}

//func TestQuickSortInt_time(t *testing.T) {
//	li := tools.Gen_ints_list(100000)
//	start := time.Now()
//	QuickSortInt(li)
//	d := time.Since(start)
//	if !isSorted(li) {
//		t.FailNow()
//	}
//	fmt.Println(d)
//
//}

func Test_quickSortInt(t *testing.T) {

}

func Test_quickSortIntPartition(t *testing.T) {

}
