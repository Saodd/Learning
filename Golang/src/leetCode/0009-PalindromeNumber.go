/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

进阶:

你能不将整数转为字符串来解决这个问题吗？


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetCode

import "fmt"

func Main0009() {
	check0009(isPalindromeNumber(121), true)
	check0009(isPalindromeNumber(-121), false)
	check0009(isPalindromeNumber(10), false)
	check0009(isPalindromeNumber(2112), true)
}

func check0009(yours bool, answer bool) {
	if yours == answer {
		fmt.Println("Pass")
	} else {
		fmt.Println("Failed!! ", "Answer:", answer, "Yours: ", yours)
	}
}

func isPalindromeNumber_Answer(x int) bool {
	// 官方答案
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
		fmt.Println(x, revertedNumber)
	}
	return x == revertedNumber || x == revertedNumber/10
}

func isPalindromeNumber(x int) bool {
	/*
		执行用时 : 20 ms, 在所有 Go 提交中击败了87.87%的用户
		内存消耗 : 6.1 MB, 在所有 Go 提交中击败了14.24%的用户
	*/
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var xlist []int
	for ; x != 0; x /= 10 {
		xlist = append(xlist, x%10)
	}

	for lo, hi := 0, len(xlist)-1; lo < hi; {
		if lo != hi {
			if xlist[lo] != xlist[hi] {
				return false
			}
		}
		lo++
		hi--
	}
	return true

}
