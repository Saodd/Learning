package p004x

import (
	"leetCode/p003x"
	"sort"
)

/*
46. 全排列

给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var le int = len(nums)
	var total int = 1
	for i := 2; i <= le; i++ {
		total *= i
	}
	var result [][]int = make([][]int, total)
	for i := 0; i < total; i++ {
		result[i] = make([]int, le)
	}

	sort.Ints(nums)
	copy(result[0], nums)
	for i := 1; i < total; i++ {
		p003x.NextPermutation(nums)
		copy(result[i], nums)
	}

	return result
}
