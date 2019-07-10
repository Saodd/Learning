/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetCode

import "fmt"

func Main0005() {
	var input string
	var output []string

	input = "babad"
	output = []string{"bab", "aba"}
	check0005(longestPalindrome(input), output)

	input = "cbbd"
	output = []string{"bb"}
	check0005(longestPalindrome(input), output)

	input = "a"
	output = []string{"a"}
	check0005(longestPalindrome(input), output)

	input = "ab"
	output = []string{"a", "b"}
	check0005(longestPalindrome(input), output)

	input = ""
	output = []string{""}
	check0005(longestPalindrome(input), output)

	input = "bb"
	output = []string{"bb"}
	check0005(longestPalindrome(input), output)

	input = "abcda"
	output = []string{"a", "b", "c", "d"}
	check0005(longestPalindrome(input), output)

	input = "babadada"
	output = []string{"adada"}
	check0005(longestPalindrome(input), output)
}

func check0005(s string, sanswer []string) {
	for _, sa := range sanswer {
		if s == sa {
			fmt.Println("Pass")
			return
		}
	}
	fmt.Println("Failed!! ", "Answer:", sanswer, "Yours: ", s)
}

func longestPalindrome(s string) string {
	lengthTol := len(s)
	lengthRL := lengthTol*2 + 1
	if lengthTol == 0 {
		return ""
	}
	// init --
	const null = byte('#')
	sb := make([]byte, lengthRL)
	sb[lengthTol*2] = null
	for i := 0; i < lengthTol*2; i += 2 {
		sb[i], sb[i+1] = null, s[i/2]
	}
	// push --
	RL := make([]int, lengthRL)
	posRight, maxRight := 0, 0
	for i := range sb {
		if i < maxRight {
			RL[i] = mymin(RL[posRight*2-i], maxRight-i)
		} else {
			RL[i] = 1
		}
		for (i-RL[i] >= 0) && (i+RL[i] < lengthRL) && (sb[i-RL[i]] == sb[i+RL[i]]) {
			RL[i]++
		}
		if i+RL[i]-1 > maxRight {
			posRight, maxRight = i, i+RL[i]-1
		}

		// print
		//for _, b := range sb {
		//	fmt.Print(string(b))
		//}
		//fmt.Println()
		//for _, v := range RL {
		//	fmt.Print(v)
		//}
		//fmt.Println()
	}
	// print
	//for _, b := range sb {
	//	fmt.Print(string(b))
	//}
	//fmt.Println()
	//for _, v := range RL {
	//	fmt.Print(v)
	//}
	//fmt.Println()
	//find max --
	maxR, maxpos := 0, 0
	for i, r := range RL {
		if r > maxR {
			maxR, maxpos = r, i
		}
	}
	var result []byte
	for _, b := range sb[(maxpos-maxR)+1 : (maxpos + maxR)] {
		if b != null {
			result = append(result, b)
		}
	}
	return string(result)
}

func longestPalindrome_mid(s string) string {
	sb := []byte(s)
	lengthTol := len(s)
	if lengthTol == 0 {
		return ""
	}
	longest := []byte{}

	for i := range sb {
		id := i * 2
		stop := id - lengthTol
		// 奇数回文
		for j := i; (j >= 0) && (j > stop); j-- {
			//fmt.Print(string(sb[j]), string(sb[id-j]))
			if sb[j] == sb[id-j] {
				//fmt.Println((id - j*2 + 1), len(longest))
				if (id - j*2 + 1) > len(longest) {
					longest = sb[j : id-j+1]
					//fmt.Println("longest 奇数", string(longest))
				}
			} else {
				break
			}
		}
		// 偶数回文
		for j := i; (j > 0) && (j > stop); j-- {
			if sb[j-1] == sb[id-j] {
				if (id - j*2 + 2) > len(longest) {
					longest = sb[j-1 : id-j+1]
					//fmt.Println("longest 偶数", string(longest))
				}
			} else {
				break
			}
		}
	}
	return string(longest)
}

func longestPalindrome_end(s string) string {
	sb := []byte(s)
	longest := []byte{}

	for i := range sb {
		for j := 0; j <= i; j++ {
			isPalindrome := true
			for k := j; k <= i-(k-j); k++ {
				if sb[k] != sb[i-(k-j)] {
					isPalindrome = false
					break
				}
			}
			if isPalindrome && (i-j+1 > len(longest)) {
				longest = sb[j : i+1]
			}
		}
	}
	return string(longest)
}
