/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetCode

import (
	"fmt"
)

func Main0007() {
	check0007(reverseInteger(123), 321)
	check0007(reverseInteger(-123), -321)
	check0007(reverseInteger(120), 21)
	check0007(reverseInteger(1563847412), 0) //2147483651
	//check0007(reverseInteger(2147483648), 0)
	//check0007(reverseInteger(2147483648), 0)
	//check0007(reverseInteger(-2147483648), 0)
}

func check0007(yours int, answer int) {
	if yours == answer {
		fmt.Println("Pass")
	} else {
		fmt.Println("Failed!! ", "Answer:", answer, "Yours: ", yours)
	}
}

func reverseInteger(x int) int {
	const limitPos int = 214748364
	const limitNeg int = -214748364
	var pop, push int
	for ; x != 0; x /= 10 {
		pop = x % 10
		if push > limitPos || (push == limitPos && pop > 7) {
			return 0
		}
		if push < limitNeg || (push == limitNeg && pop < -8) {
			return 0
		}
		push = push*10 + pop
	}
	return push
}

func reverseInteger_Brute(x int) int {
	if x < 10 && x > -10 {
		return x
	}

	neg := false
	if x < 0 {
		neg = true
		x *= -1
	}

	var result []int
	for temp := 0; x > 0; {
		temp = x - x/10*10
		result = append(result, temp)
		x /= 10
	}

	length := len(result)
	var resultInt int
	for i, b := range result[1:] {
		resultInt += b * myPow10(length-i-2)
	}

	//fmt.Println(resultInt, length, result[9])
	if length > 10 {
		return 0
	}
	if length == 10 {
		if result[0] > 2 {
			return 0
		} else if result[0] == 2 {
			if neg {
				if resultInt > 147483648 {
					return 0
				}
			} else {
				if resultInt > 147483647 {
					return 0
				}
			}
		}
	}

	if neg {
		return (result[0]*myPow10(length-1) + resultInt) * -1
	}
	return result[0]*myPow10(length-1) + resultInt
}

func myPow10(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	case 4:
		return 10000
	case 5:
		return 100000
	case 6:
		return 1000000
	case 7:
		return 10000000
	case 8:
		return 100000000
	case 9:
		return 1000000000
	case 10:
		return 10000000000
	}
	return 0
}
