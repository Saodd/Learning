package p003x

import (
    "reflect"
    "testing"
)

//func Test_searchRangeLow(t *testing.T) {
//    type args struct {
//        nums   []int
//        target int
//    }
//    tests := []struct {
//        name string
//        args args
//        want int
//    }{
//        {
//            args: args{[]int{1, 2, 3, 3, 3, 3, 3}, 3},
//            want: 2,
//        },
//        {
//            args: args{[]int{1, 2, 3, 3, 3, 3, 3}, 4},
//            want: -1,
//        },
//        {
//            args: args{[]int{1, 2, 3, 4}, 4},
//            want: 3,
//        },
//    }
//    for _, tt := range tests {
//        t.Run(tt.name, func(t *testing.T) {
//            if got := searchRangeLow(tt.args.nums, tt.args.target); got != tt.want {
//                t.Errorf("searchRangeLow() = %v, want %v", got, tt.want)
//            }
//        })
//    }
//}
//
//func Test_searchRangeHigh(t *testing.T) {
//    type args struct {
//        nums   []int
//        target int
//    }
//    tests := []struct {
//        name string
//        args args
//        want int
//    }{
//        {
//            args: args{[]int{1, 2, 3, 4}, 1},
//            want: 0,
//        },
//        {
//            args: args{[]int{1, 1, 1, 2, 3, 4}, 1},
//            want: 2,
//        },
//        {
//            name: "miss left",
//            args: args{[]int{1, 2, 3, 4}, 0},
//            want: -1,
//        },
//        {
//            name: "miss right",
//            args: args{[]int{1, 2, 3, 4}, 5},
//            want: -1,
//        },
//    }
//    for _, tt := range tests {
//        t.Run(tt.name, func(t *testing.T) {
//            if got := searchRangeHigh(tt.args.nums, tt.args.target); got != tt.want {
//                t.Errorf("searchRangeHigh() = %v, want %v", got, tt.want)
//            }
//        })
//    }
//}

func Test_searchRange(t *testing.T) {
    type args struct {
        nums   []int
        target int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            args: args{[]int{5, 7, 7, 8, 8, 10}, 8},
            want: []int{3, 4},
        },
        {
            args: args{[]int{5, 7, 7, 8, 8, 10}, 6},
            want: []int{-1, -1},
        },
        {
            args: args{[]int{2, 2}, 2},
            want: []int{0, 1},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("searchRange() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_searchRangeLeft(t *testing.T) {
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
            args: args{[]int{1, 2, 4, 5}, 2},
            want: 1,
        },
        {
            args: args{[]int{1, 2, 2, 4, 5}, 2},
            want: 1,
        },
        {
            name:"边缘",
            args: args{[]int{1, 1, 2, 4, 5}, 1},
            want: 0,
        },
        {
            name:"边缘",
            args: args{[]int{1, 2, 2, 5, 5}, 5},
            want: 3,
        },

        {
            name: "不存在",
            args: args{[]int{1, 2, 4, 5}, 3},
            want: -1,
        },
        {
            name: "不存在",
            args: args{[]int{1, 2, 4, 5}, 0},
            want: -1,
        }, {
            name: "不存在",
            args: args{[]int{1, 2, 4, 5}, 6},
            want: -1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := searchRangeLeft(tt.args.nums, tt.args.target); got != tt.want {
                t.Errorf("searchRangeLeft() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_searchRangeRight(t *testing.T) {
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
            args: args{[]int{1, 2, 4, 5}, 2},
            want: 1,
        },
        {
            args: args{[]int{1, 2, 2, 4, 5}, 2},
            want: 2,
        },
        {
            name:"边缘",
            args: args{[]int{1, 1, 2, 4, 5}, 1},
            want: 1,
        },
        {
            name:"边缘",
            args: args{[]int{1, 2, 2, 5, 5}, 5},
            want: 4,
        },

        {
            name: "不存在",
            args: args{[]int{1, 2, 4, 5}, 3},
            want: -1,
        },
        {
            name: "不存在",
            args: args{[]int{1, 2, 4, 5}, 0},
            want: -1,
        }, {
            name: "不存在",
            args: args{[]int{1, 2, 4, 5}, 6},
            want: -1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := searchRangeRight(tt.args.nums, tt.args.target); got != tt.want {
                t.Errorf("searchRangeRight() = %v, want %v", got, tt.want)
            }
        })
    }
}
