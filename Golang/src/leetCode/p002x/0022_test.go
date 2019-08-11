package p002x

import (
    "reflect"
    "testing"
)

func Test_generateParenthesis(t *testing.T) {
    type args struct {
        n int
    }
    tests := []struct {
        name string
        args args
        want []string
    }{
        {
            args: args{0},
            want: []string{},
        },
        {
            args: args{1},
            want: []string{"()"},
        },
        {
            args: args{2},
            want: []string{"(())","()()"},
        },
        {
            args: args{3},
            want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
            }
        })
    }
}
