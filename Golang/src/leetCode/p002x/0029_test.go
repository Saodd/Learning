package p002x

import "testing"

func Test_divide(t *testing.T) {
    const maxInt = 1<<31 - 1
    const minInt = -1 << 31
    type args struct {
        dividend int
        divisor  int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            args: args{4, -2},
            want: -2,
        },
        {
            args: args{-4, -2},
            want: 2,
        },
        {
            args: args{-4, 2},
            want: -2,
        },
        {
            args: args{4, 2},
            want: 2,
        },
        {
            args: args{1235154145, 1234},
            want: 1235154145 / 1234,
        },
        {
            args: args{maxInt, maxInt},
            want: maxInt / maxInt,
        },
        {
            args: args{minInt, maxInt},
            want: minInt / maxInt,
        },
        {
            args: args{maxInt, minInt},
            want: maxInt / minInt,
        },
        {
            args: args{minInt, minInt},
            want: minInt / minInt,
        },
        {
            args: args{minInt, -1},
            want: maxInt, // 边界特例
        },
        {
            args: args{maxInt - 10, 1 << 30},
            want: (maxInt - 10) / (1 << 30),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
                t.Errorf("divide() = %v, want %v", got, tt.want)
            }
        })
    }
}
