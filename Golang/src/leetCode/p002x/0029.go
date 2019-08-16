package p002x

/*
给定两个整数，被除数 dividend 和除数 divisor。
将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3

示例 2:

输入: dividend = 7, divisor = -3
输出: -2

说明:

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是
[−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/divide-two-integers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func divide(dividend int, divisor int) (result int) {
    // 是否异号
    sign := dividend ^ divisor

    // 边界溢出检查
    if dividend == minInt {
        if divisor >= 0 {
            dividend += divisor
            defer func() { result -= 1 }()
        } else {
            dividend -= divisor
            defer func() {
                if result != maxInt {
                    result += 1
                }
            }()
        }
    }
    if divisor == minInt {
        return 0
    }

    // 都转换为正数
    dividend = (dividend ^ (dividend >> 31)) - (dividend >> 31)
    divisor = (divisor ^ (divisor >> 31)) - (divisor >> 31)

    // 找出结果的位数，防止左移溢出
    var dig uint
    for dig = 1; dividend >= divisor<<dig; dig++ {
        if divisor>>(30-dig) != 0 {
            dig++
            break
        }
    }
    // 开始计算
    for i := dig - 1; i > 0; i-- {
        if dividend >= divisor<<i {
            dividend -= divisor << i
            result |= 1 << i
        }
    }
    if dividend >= divisor {
        result |= 1
    }
    // 最后的正负处理
    if sign >= 0 {
        return result
    }
    return ^result + 1
}

const maxInt = 1<<31 - 1
const minInt = -1 << 31
