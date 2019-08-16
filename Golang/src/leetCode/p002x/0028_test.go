package p002x

import (
    "fmt"
    "math/rand"
    "strings"
    "testing"
    "time"
    "unsafe"
)

func Test_strStr(t *testing.T) {
    type args struct {
        haystack string
        needle   string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            args: args{"hello", "ll"},
            want: 2,
        },
        {
            args: args{"aaaaa", "bba"},
            want: -1,
        },
        {
            args: args{"hello", "hello"},
            want: 0,
        },
        {
            args: args{"hello", "loo"},
            want: -1,
        },
        {
            args: args{"hello", "helloo"},
            want: -1,
        },
        {
            args: args{"hello", ""},
            want: 0,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
                t.Errorf("strStr() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_strStr_Time(t *testing.T) {
    var haystack = make([]byte, 100000000)
    var needle = make([]byte, 100)
    rand.Read(haystack)
    rand.Read(needle)
    var s, s2 string = *(*string)(unsafe.Pointer(&haystack)), *(*string)(unsafe.Pointer(&needle))
    {
       start := time.Now()
       got := strStr(s, s2)
       totalTime := time.Since(start)
       fmt.Println(got, totalTime.Seconds())
    }
    {
       start := time.Now()
       got := strings.Index(s, s2)
       totalTime := time.Since(start)
       fmt.Println(got, totalTime.Seconds())
    }
}
