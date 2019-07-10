/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/

package leetCode

import (
	"fmt"
)

func Main0003() {
	instances := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
		"":         0,
		" ":        1,
		"au":       2,
		"aab":      2,
		"dvdf":     3,
	}
	for k, v := range instances {
		fmt.Printf("Input: '%s', Answer: %v, Yours: %v \n", k, v, lengthOfLongestSubstring(k))
	}

}

func lengthOfLongestSubstring(s string) int {
	begin, maxL := 0, 0
	var p, j int
	for p = 1; p < len(s); p++ {
		b := s[p]
		// check isIn
		for j = begin; j < p; j++ {
			if b == s[j] {
				if p-begin > maxL {
					maxL = p - begin
				}
				begin = j + 1
				break
			}
		}
	}
	if len(s)-begin > maxL {
		maxL = len(s) - begin
	}
	return maxL
}

/*
后记：
一开始没有想到好的思路，提交失败了好多次……然后才想到这个办法。
时间复杂度不会太高，最坏情况下应该是(n*字符种类数量)，也就是说如果全英文字符的话，最多(n*26)的时间复杂度。

我的实现在网站上给出的成绩是：
执行用时 :4 ms， 在所有 Go 提交中击败了96.22%的用户
内存消耗 :2.6 MB, 在所有 Go 提交中击败了89.69%的用户

然后对比了一下评论里说的：

C++
执行用时 : 16 ms, 在Longest Substring Without Repeating Characters的C++提交中击败了99.48% 的用户
内存消耗 : 9.3 MB, 在Longest Substring Without Repeating Characters的C++提交中击败了87.06% 的用户

Python3:
执行用时 : 64 ms , 在所有 Python3 提交中击败了 99.80% 的用户
内存消耗 : 13.2 MB , 在所有 Python3 提交中击败了 76.53% 的用户

发现Golang比C++省内存？
*/
