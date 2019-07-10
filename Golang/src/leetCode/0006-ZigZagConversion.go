/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

示例 1:
输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"

示例 2:
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:
L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetCode

import (
	"fmt"
)

func Main0006() {
	check0006(convertZigZag("ABCDEFG", 3), "AEBDFCG")
	check0006(convertZigZag("LEETCODEISHIRING", 3), "LCIRETOESIIGEDHN")
	check0006(convertZigZag("LEETCODEISHIRING", 4), "LDREOEIIECIHNTSG")
	check0006(convertZigZag("AB", 1), "AB")
	check0006(convertZigZag("A", 2), "A")
	check0006(convertZigZag("", 2), "")
	check0006(convertZigZag("ABCDE", 4), "ABCED")
	check0006(convertZigZag("ABC", 3), "ABC")

}

func check0006(s string, sanswer string) {
	if s == sanswer {
		fmt.Println("Pass")
	} else {
		fmt.Println("Failed!! ", "Answer:", sanswer, "Yours: ", s)
	}
}

func convertZigZag(s string, numRows int) string {
	lengthS := len(s)
	if numRows == 1 {
		return s
	}
	var result []byte

	for n := 0; n < numRows; n++ {
		for c := 0; c < (lengthS/(2*numRows-2) + 2); c++ {
			ileft := c*2*(numRows-1) - n
			iright := ileft + 2*n
			//fmt.Println(n, c,(2*numRows-2) , ileft, iright)
			if ileft >= lengthS {
				break
			}
			if ileft == iright {
				//fmt.Println(string(s[ileft]))
				result = append(result, s[ileft])
			} else {
				if ileft >= 0 && (n != numRows-1) {
					//fmt.Println(string(s[ileft]))
					result = append(result, s[ileft])
				}
				if iright < lengthS {
					//fmt.Println(string(s[iright]))
					result = append(result, s[iright])
				}
			}
		}
	}
	return string(result)
}

func convertZigZag_brute(s string, numRows int) string {
	result := make([][]byte, numRows)
	lengthS := len(s)
	for i := 0; i < lengthS; {
		for pos := 0; pos < numRows && i < lengthS; pos++ {
			result[pos] = append(result[pos], s[i])
			i++
		}
		for neg := numRows - 2; neg > 0 && i < lengthS; neg-- {
			result[neg] = append(result[neg], s[i])
			i++
		}
	}
	var resultS string
	for _, s := range result {
		resultS += string(s)
	}
	return resultS
}

func convertZigZag_try1(s string, numRows int) string {
	if len(s) == 0 {
		return ""
	}
	result := make([][]byte, numRows)
	//result[0] = append(result[0], s[0])
	direction := false
	numSlice := numRows - 1
	for r, i := 0, 0; i < len(s); i++ {
		if r == 0 || r == numSlice {
			direction = !direction
		}
		fmt.Println(r, i)
		result[r] = append(result[r], s[i])
		fmt.Println(r, string(result[r]))
		if direction {
			r++
		} else {
			r--
		}
	}
	var resultS string
	for _, s := range result {
		resultS += string(s)
	}
	return resultS
}
