package p004x

import "testing"

func Test_multiply(t *testing.T) {
    type args struct {
        num1 string
        num2 string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "示例",
            args: args{"2", "3"},
            want: "6",
        },
        {
            name: "示例",
            args: args{"123", "456"},
            want: "56088",
        },
        {
            name: "零",
            args: args{"0", "456"},
            want: "0",
        },
        {
            name: "大数",
            args: args{"11111111111111111111111111111111111111111111111111", "22"},
            want: "244444444444444444444444444444444444444444444444442",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
                t.Errorf("multiply() = %v, want %v", got, tt.want)
            }
        })
    }
}
