package p001x

import (
	"fmt"
	"testing"
	"time"
)

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{1},
			want: "I",
		},
		{
			args: args{3},
			want: "III",
		},
		{
			args: args{4},
			want: "IV",
		},
		{
			args: args{9},
			want: "IX",
		},
		{
			args: args{58},
			want: "LVIII",
		},
		{
			args: args{1994},
			want: "MCMXCIV",
		},
		{
			args: args{10},
			want: "X",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_intToRoman_Time(t *testing.T) {
	var totalTime time.Duration
	var loopNumber int = 10000
	start := time.Now()
	for i := 0; i < loopNumber; i++ {
		for p:=1; p<4000; p++{
			//intToRoman(p)
			nothing(p)
		}
	}
	totalTime = time.Since(start)
	fmt.Println(totalTime.Seconds()/float64(loopNumber))
}

func nothing(p int)  {
	p++
}