package p001x

import "unsafe"

/*
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，
即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如
 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大
 数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的
 规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

示例 1:
输入: 3
输出: "III"

示例 2:
输入: 4
输出: "IV"

示例 3:
输入: 9
输出: "IX"

示例 4:
输入: 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.

示例 5:
输入: 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

var RomanTable = [7]byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
func intToRoman(num int) string {
	// 题目规定 1<= num <=3999
	var r int
	var result = make([]byte, 16)
	var p int = 15
	for i := 0; i < 4; i++ {
		r = i*2
		switch num % 10 {
		case 1:
			result[p] = RomanTable[r]
			p--
		case 2:
			result[p-1], result[p] = RomanTable[r], RomanTable[r]
			p -= 2
		case 3:
			result[p-2], result[p-1], result[p] = RomanTable[r], RomanTable[r], RomanTable[r]
			p -= 3
		case 4:
			result[p-1], result[p] = RomanTable[r], RomanTable[r+1]
			p -= 2
		case 5:
			result[p] = RomanTable[r+1]
			p--
		case 6:
			result[p-1], result[p] = RomanTable[r+1], RomanTable[r]
			p -= 2
		case 7:
			result[p-2], result[p-1], result[p] = RomanTable[r+1], RomanTable[r], RomanTable[r]
			p -= 3
		case 8:
			result[p-3], result[p-2], result[p-1], result[p] = RomanTable[r+1], RomanTable[r], RomanTable[r], RomanTable[r]
			p -= 4
		case 9:
			result[p-1], result[p] = RomanTable[r], RomanTable[r+2]
			p -= 2
		}
		num /= 10
	}
	result = result[p+1:]
	return (*(*string)(unsafe.Pointer(&(result))))
}


//var RomanTable = [7]byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
//func intToRoman(num int) string {
//	// 题目规定 1<= num <=3999
//	var romanTable []byte
//	var result = make([]byte, 16)
//	var p int = 15
//	for i := 0; i < 4; i++ {
//		romanTable = RomanTable[i*2:]
//		switch num % 10 {
//		case 1:
//			result[p] = romanTable[0]
//			p--
//		case 2:
//			result[p-1], result[p] = romanTable[0], romanTable[0]
//			p -= 2
//		case 3:
//			result[p-2], result[p-1], result[p] = romanTable[0], romanTable[0], romanTable[0]
//			p -= 3
//		case 4:
//			result[p-1], result[p] = romanTable[0], romanTable[1]
//			p -= 2
//		case 5:
//			result[p] = romanTable[1]
//			p--
//		case 6:
//			result[p-1], result[p] = romanTable[1], romanTable[0]
//			p -= 2
//		case 7:
//			result[p-2], result[p-1], result[p] = romanTable[1], romanTable[0], romanTable[0]
//			p -= 3
//		case 8:
//			result[p-3], result[p-2], result[p-1], result[p] = romanTable[1], romanTable[0], romanTable[0], romanTable[0]
//			p -= 4
//		case 9:
//			result[p-1], result[p] = romanTable[0], romanTable[2]
//			p -= 2
//		}
//		num /= 10
//	}
//	return (*(*string)(unsafe.Pointer(&(result))))[p+1:]
//}



//var RomanTable = [7]string{"I", "V", "X", "L", "C", "D", "M"}
//func intToRoman(num int) string {
//	// 题目规定 1<= num <=3999
//	var romanTable []string
//	var result = make([]string, 4)
//	for i := 0; i < 4; i++ {
//		romanTable = RomanTable[i*2:]
//		switch num % 10 {
//		case 1:
//			result[3-i] = fmt.Sprint(romanTable[0])
//		case 2:
//			result[3-i] = fmt.Sprint(romanTable[0], romanTable[0])
//		case 3:
//			result[3-i] = fmt.Sprint(romanTable[0], romanTable[0], romanTable[0])
//		case 4:
//			result[3-i] = fmt.Sprint(romanTable[0], romanTable[1])
//		case 5:
//			result[3-i] = fmt.Sprint(romanTable[1])
//		case 6:
//			result[3-i] = fmt.Sprint(romanTable[1], romanTable[0])
//		case 7:
//			result[3-i] = fmt.Sprint(romanTable[1], romanTable[0], romanTable[0])
//		case 8:
//			result[3-i] = fmt.Sprint(romanTable[1], romanTable[0], romanTable[0], romanTable[0])
//		case 9:
//			result[3-i] = fmt.Sprint(romanTable[0], romanTable[2])
//		}
//		num /= 10
//	}
//	return strings.Join(result, "")
//}
