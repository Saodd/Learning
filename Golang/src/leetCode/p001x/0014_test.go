package p001x

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{[]string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			args: args{[]string{"dog","racecar","car"}},
			want: "",
		},
		{
			args: args{[]string{}},
			want: "",
		},
		{
			args: args{[]string{"haha", ""}},
			want: "",
		},
		{
			args: args{[]string{""}},
			want: "",
		},
		{
			args: args{[]string{"a"}},
			want: "a",
		},
		{
			args: args{[]string{"c", "c"}},
			want: "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
