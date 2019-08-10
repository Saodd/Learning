package p002x

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
func isValid(s string) bool {
    if len(s)%2 != 0 {
        return false
    }
    var halfLen, p int = len(s)/2, 0
    var stack []byte = make([]byte, halfLen)
    for i, le := 0, len(s); i < le; i++ {
        c := s[i]
        switch c {
        case '(', '[', '{':
            if p == halfLen {
                return false
            }
            stack[p] = c
            p++
        case ')', ']', '}':
            if p == 0 {
                return false
            }
            if c-stack[p-1] <= 2 { // 40 41 91 93 123 125
                p--
            } else {
                return false
            }
        }
    }
    if p==0{
        return true
    }
    return false
}

//func isValid(s string) bool {
//    var a, b, c int = 0, 0, 0
//    for _, i := range s {
//        switch i {
//        case '(':
//            a++
//        case '[':
//            b++
//        case '{':
//            c++
//        case ')':
//            a--
//        case ']':
//            b--
//        case '}':
//            c--
//        }
//    }
//    if a == 0 && b == 0 && c == 0 {
//        return true
//    }
//    return false
//}
