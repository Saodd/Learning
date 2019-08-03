package p001x

import (
	"fmt"
	"testing"
	"time"
)

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{"I"},
			want: 1,
		},
		{
			args: args{"III"},
			want: 3,
		},
		{
			args: args{"IV"},
			want: 4,
		},
		{
			args: args{"IX"},
			want: 9,
		},
		{
			args: args{"LVIII"},
			want: 58,
		},
		{
			args: args{"MCMXCIV"},
			want: 1994,
		},
		{
			args: args{"X"},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_romanToInt_Time(t *testing.T) {
	// init -------
	var input = make([]string, 4000)
	for i:=1; i<4000; i++{
		input[i] = intToRoman(i)
	}
	// test --------
	var totalTime time.Duration
	var loopNumber int = 10000
	start := time.Now()
	for i := 0; i < loopNumber; i++ {
		for p:=1; p<4000; p++{
			romanToInt(input[p])
		}
	}
	totalTime = time.Since(start)
	fmt.Println(totalTime.Seconds()/float64(loopNumber))
}