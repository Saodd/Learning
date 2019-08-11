package p002x

/*
给出 n 代表生成括号的对数，
请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func generateParenthesis(n int) []string {
    result0022 = make([]string, 0)
    if n == 0 {
        return result0022
    }
    length0022 = 2 * n
    temp0022 = make([]byte, length0022)
    recGenParenthesis(n, 0, 0)
    return result0022
}

var result0022 []string
var temp0022 []byte
var length0022 int

func recGenParenthesis(n, pos, count int) {
    if n == 0 { // 左括号用完了，后面补全右括号
        for i := pos; i < length0022; i++ {
            temp0022[i] = ')'
        }
        result0022 = append(result0022, string(temp0022))
        return
    }
    // 还有左括号，那么分两种情况：1.放一个左括号
    temp0022[pos] = '('
    recGenParenthesis(n-1, pos+1, count+1)
    // 2.放一个右括号
    if count > 0 {
        temp0022[pos] = ')'
        recGenParenthesis(n, pos+1, count-1)
    }
}
