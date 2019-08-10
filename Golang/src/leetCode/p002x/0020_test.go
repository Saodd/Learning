/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package p002x

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"1",
			args:args{"()"},
			want:true,
		},
		{
			name:"2",
			args:args{"()[]{}"},
			want:true,
		},
		{
			name:"3",
			args:args{"(]"},
			want:false,
		},
		{
			name:"4",
			args:args{"([)]"},
			want:false,
		},
		{
			name:"5",
			args:args{"{[]}"},
			want:true,
		},
		{
			name:"6",
			args:args{"(("},
			want:false,
		},
		{
			name:"7",
			args:args{"))"},
			want:false,
		},
		{
			name:"8",
			args:args{"(()("},
			want:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
