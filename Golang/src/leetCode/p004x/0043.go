package p004x

import (
    "unsafe"
)

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，
返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"

示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"

说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/multiply-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func multiply(num1 string, num2 string) string {
    if num1[0] == '0' || num2[0] == '0' {
        return "0"
    }
    // 初始化
    var result = make([]int32, len(num1)+len(num2))
    var n1 = make([]int32, len(num1))
    var n2 = make([]int32, len(num2))
    {

        for i, le := 0, len(num1)-1; i <= le; i++ {
            n1[le-i] = int32(num1[i] - '0')
        }
        for i, le := 0, len(num2)-1; i <= le; i++ {
            n2[le-i] = int32(num2[i] - '0')
        }
    }
    // 逐位计算
    var incre int32 = 0
    for i1, dig1 := range n1 {
        for i2, dig2 := range n2 {
            incre += (dig1*dig2 + result[i1+i2])
            result[i1+i2] = incre % 10
            incre /= 10
        }
        result[i1+len(num2)] = incre
        incre = 0
    }
    // 返回结果
    var resultS = make([]byte, len(num1)+len(num2))
    for i, le := 0, len(result)-1; i <= le; i++ {
        resultS[le-i] = byte(result[i]) + '0'
    }
    for i, b := range resultS {
        if b != '0' {
            resultS = resultS[i:]
            break
        }
    }
    return *(*string)(unsafe.Pointer(&resultS))
}
