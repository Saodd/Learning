/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regular-expression-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package p0010_RegularExpression

import "fmt"

var result [][]int

func isMatch(s string, p string) bool {
	result = make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		for j := 0; j < len(p)+1; j++ {
			result[i] = append(result[i], 0)
		}
	}

	return subMatch(0, 0, []byte(s), []byte(p))
}

func subMatch(x, y int, s, p []byte) bool {
	if result[x][y] != 0 {
		if result[x][y] == 1 {
			return true
		}
		return false
	}
	var ans bool
	if y >= len(p) {
		ans = (x >= len(s))
	} else {
		//fmt.Println(x, len(p))
		first_match := (x < len(s)) && (s[x] == p[y] || p[y] == '.')
		if (y+1 < len(p)) && (p[y+1] == '*') {
			ans = subMatch(x, y+2, s, p) || (first_match && subMatch(x+1, y, s, p))
		} else {
			ans = first_match && subMatch(x+1, y+1, s, p)
		}
	}
	if ans {
		result[x][y] = 1
	} else {
		result[x][y] = -1
	}
	return ans
}

//-------------------------------------------------------------
func Main0010() {
	check("aa", "a", false)
	check("aa", "a*", true)
	check("ab", ".*", true)
	check("aab", "c*a*b", true)
	check("mississippi", "mis*is*p*.", false)
	//result = make([][]int, 10,10)
	//isMatch("a", "a")
	//fmt.Println(result[1][1])

}

func check(s, p string, answer bool) {
	yours := isMatch(s, p)
	if yours == answer {
		fmt.Println("Pass")
	} else {
		fmt.Println("Failed!! ", "Answer:", answer, "Yours: ", yours)
	}
}
